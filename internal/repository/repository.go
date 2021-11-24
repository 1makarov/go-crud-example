package repository

import (
	"context"
	"github.com/1makarov/go-crud-example/internal/repository/books"
	"github.com/1makarov/go-crud-example/internal/types"
	"github.com/jmoiron/sqlx"
)

type Books interface {
	Create(ctx context.Context, v types.BookCreateInput) error
	GetByID(ctx context.Context, id int) (*types.Book, error)
	GetAll(ctx context.Context) ([]types.Book, error)
	DeleteByID(ctx context.Context, id int) error
	UpdateByID(ctx context.Context, id int, v types.BookUpdateInput) error
}

type Repository struct {
	Books
}

func New(db *sqlx.DB) *Repository {
	return &Repository{
		Books: books.NewRepoBooks(db),
	}
}
