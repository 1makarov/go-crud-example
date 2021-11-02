package repository

import (
	"context"
	"github.com/1makarov/go-crud-example/internal/model"
	"github.com/1makarov/go-crud-example/internal/repository/books"
	"github.com/jmoiron/sqlx"
)

type Books interface {
	Create(ctx context.Context, v model.BookCreateInput) error
	GetByID(ctx context.Context, id int) (*model.Book, error)
	GetAll(ctx context.Context) ([]model.Book, error)
	DeleteByID(ctx context.Context, id int) error
	UpdateByID(ctx context.Context, id int, v model.BookUpdateInput) error
}

type Repository struct {
	Books
}

func New(db *sqlx.DB) *Repository {
	return &Repository{
		Books: books.NewRepo(db),
	}
}
