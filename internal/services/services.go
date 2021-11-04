package services

import (
	"context"
	authorization "github.com/1makarov/go-crud-example/internal/pkg/auth"
	"github.com/1makarov/go-crud-example/internal/repository"
	"github.com/1makarov/go-crud-example/internal/services/auth"
	"github.com/1makarov/go-crud-example/internal/services/books"
	"github.com/1makarov/go-crud-example/internal/types"
)

type Books interface {
	Create(ctx context.Context, v types.BookCreateInput) error
	GetByID(ctx context.Context, id int) (*types.Book, error)
	GetAll(ctx context.Context) ([]types.Book, error)
	DeleteByID(ctx context.Context, id int) error
	UpdateByID(ctx context.Context, id int, v types.BookUpdateInput) error
}

type Auth interface {
	CreateToken() (string, error)
	ValidToken(token string) error
}

type Service struct {
	Books
	Auth
}

func New(repo *repository.Repository, a *authorization.Auth) *Service {
	return &Service{
		Books: books.Init(repo.Books),
		Auth:  auth.Init(a.JWT),
	}
}
