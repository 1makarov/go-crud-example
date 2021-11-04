package books

import (
	"context"
	"github.com/1makarov/go-crud-example/internal/repository"
	"github.com/1makarov/go-crud-example/internal/types"
)

type Books struct {
	repo repository.Books
}

func Init(repo repository.Books) *Books {
	return &Books{
		repo: repo,
	}
}

func (b *Books) Create(ctx context.Context, v types.BookCreateInput) error {
	return b.repo.Create(ctx, v)
}

func (b *Books) GetByID(ctx context.Context, id int) (*types.Book, error) {
	return b.repo.GetByID(ctx, id)
}

func (b *Books) GetAll(ctx context.Context) ([]types.Book, error) {
	return b.repo.GetAll(ctx)
}

func (b *Books) DeleteByID(ctx context.Context, id int) error {
	return b.repo.DeleteByID(ctx, id)
}

func (b *Books) UpdateByID(ctx context.Context, id int, v types.BookUpdateInput) error {
	return b.repo.UpdateByID(ctx, id, v)
}
