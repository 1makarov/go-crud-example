package services

import (
	"context"
	"github.com/1makarov/go-cache"
	"github.com/1makarov/go-crud-example/internal/repository"
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

type Service struct {
	Books
}

func New(repo *repository.Repository, cache *cache.Cache) *Service {
	return &Service{
		Books: books.InitServiceBooks(repo.Books, cache),
	}
}
