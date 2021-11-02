package services

import (
	"context"
	"github.com/1makarov/go-crud-example/internal/model"
	"github.com/1makarov/go-crud-example/internal/repository"
	"github.com/1makarov/go-crud-example/internal/services/books"
)

type Books interface {
	Create(ctx context.Context, v model.BookCreateInput) error
	GetByID(ctx context.Context, id int) (*model.Book, error)
	GetAll(ctx context.Context) ([]model.Book, error)
	DeleteByID(ctx context.Context, id int) error
	UpdateByID(ctx context.Context, id int, v model.BookUpdateInput) error
}

type Service struct {
	Books
}

func New(repo *repository.Repository) *Service {
	return &Service{
		Books: books.Init(repo),
	}
}
