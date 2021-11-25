package services

import (
	"context"
	"github.com/1makarov/go-cache"
	"github.com/1makarov/go-crud-example/internal/pkg/hash"
	"github.com/1makarov/go-crud-example/internal/repository"
	"github.com/1makarov/go-crud-example/internal/services/books"
	"github.com/1makarov/go-crud-example/internal/services/users"
	"github.com/1makarov/go-crud-example/internal/types"
)

type Books interface {
	Create(ctx context.Context, v types.BookCreateInput) error
	GetByID(ctx context.Context, id int) (*types.Book, error)
	GetAll(ctx context.Context) ([]types.Book, error)
	DeleteByID(ctx context.Context, id int) error
	UpdateByID(ctx context.Context, id int, v types.BookUpdateInput) error
}

type Users interface {
	SignUp(ctx context.Context, input types.SignUpInput) error
	SignIn(ctx context.Context, input types.SignInInput) (*types.User, error)
}

type Service struct {
	Books
	Users
}

func New(repo *repository.Repository, cache *cache.Cache, hash *hash.Manager) *Service {
	return &Service{
		Books: books.InitServiceBooks(repo.Books, cache),
		Users: users.InitServiceUsers(repo, hash),
	}
}
