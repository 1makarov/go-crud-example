package books

import (
	"context"
	"github.com/1makarov/go-crud-example/internal/model"
	"github.com/1makarov/go-crud-example/internal/repository"
)

type Books struct {
	repo repository.Books
}

func Init(repo repository.Books) *Books {
	return &Books{
		repo: repo,
	}
}

func (b *Books) Create(ctx context.Context, v model.BookCreateInput) error {
	return b.repo.Create(ctx, v)
}

func (b *Books) GetByID(ctx context.Context, id int) (*model.Book, error) {
	return b.repo.GetByID(ctx, id)
}

func (b *Books) GetAll(ctx context.Context) ([]model.Book, error) {
	return b.repo.GetAll(ctx)
}

func (b *Books) DeleteByID(ctx context.Context, id int) error {
	return b.repo.DeleteByID(ctx, id)
}

func (b *Books) UpdateByID(ctx context.Context, id int, v model.BookUpdateInput) error {
	return b.repo.UpdateByID(ctx, id, v)
}
