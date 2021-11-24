package books

import (
	"context"
	"github.com/1makarov/go-cache"
	"github.com/1makarov/go-crud-example/internal/repository"
	"github.com/1makarov/go-crud-example/internal/types"
)

type ServiceBooks struct {
	repo  repository.Books
	cache *cache.Cache
}

func InitServiceBooks(repo repository.Books, cache *cache.Cache) *ServiceBooks {
	return &ServiceBooks{
		repo:  repo,
		cache: cache,
	}
}

func (b *ServiceBooks) Create(ctx context.Context, v types.BookCreateInput) error {
	return b.repo.Create(ctx, v)
}

func (b *ServiceBooks) GetByID(ctx context.Context, id int) (*types.Book, error) {
	if value, err := b.cache.Get(id); err == nil {
		return value.(*types.Book), nil
	}

	book, err := b.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return book, b.cache.Set(id, book)
}

func (b *ServiceBooks) GetAll(ctx context.Context) ([]types.Book, error) {
	return b.repo.GetAll(ctx)
}

func (b *ServiceBooks) DeleteByID(ctx context.Context, id int) error {
	if err := b.repo.DeleteByID(ctx, id); err != nil {
		return err
	}

	b.cache.Delete(id)

	return nil
}

func (b *ServiceBooks) UpdateByID(ctx context.Context, id int, v types.BookUpdateInput) error {
	if err := b.repo.UpdateByID(ctx, id, v); err != nil {
		return err
	}

	if value, err := b.cache.GetAndDelete(id); err == nil {
		book := value.(*types.Book)

		if v.Name != nil {
			book.Name = *v.Name
		}

		if v.ISBN != nil {
			book.ISBN = *v.ISBN
		}

		if v.Title != nil {
			book.Title = *v.Title
		}

		if v.Description != nil {
			book.Description = *v.Description
		}

		if err = b.cache.Set(id, book); err != nil {
			return err
		}
	}

	return nil
}
