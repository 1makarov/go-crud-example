package books

import (
	"context"
	"fmt"
	"github.com/1makarov/go-crud-example/internal/model"
	"github.com/jmoiron/sqlx"
)

type Books struct {
	db *sqlx.DB
}

func NewRepo(db *sqlx.DB) *Books {
	return &Books{
		db: db,
	}
}

func (b *Books) Create(ctx context.Context, v model.BookCreateInput) error {
	if _, err := b.db.NamedExecContext(ctx, `

	insert into books (name, title, isbn, description) values (:name, :title, :isbn, :description)

	`, v); err != nil {
		return err
	}

	return nil
}

func (b *Books) GetByID(ctx context.Context, id int) (*model.Book, error) {
	var v model.Book

	if err := b.db.GetContext(ctx, &v, `select * from books where id = $1`, id); err != nil {
		return nil, err
	}

	return &v, nil
}

func (b *Books) GetAll(ctx context.Context) ([]model.Book, error) {
	var v []model.Book

	if err := b.db.SelectContext(ctx, &v, `select * from books`); err != nil {
		return nil, err
	}

	return v, nil
}

func (b *Books) DeleteByID(ctx context.Context, id int) error {
	if _, err := b.db.ExecContext(ctx, `delete from books where id = $1`, id); err != nil {
		return err
	}

	return nil
}

func (b *Books) UpdateByID(ctx context.Context, id int, v model.BookUpdateInput) error {
	var query string

	if v.Description != nil {
		query += fmt.Sprintf("description = '%s',", *v.Description)
	}
	if v.ISBN != nil {
		query += fmt.Sprintf("isbn = '%s',", *v.ISBN)
	}
	if v.Name != nil {
		query += fmt.Sprintf("name = '%s',", *v.Name)
	}
	if v.Title != nil {
		query += fmt.Sprintf("title = '%s' ", *v.Title)
	}

	r := fmt.Sprintf("update book set %s where id = %d", query[:len(query)-1], id)

	if _, err := b.db.ExecContext(ctx, r); err != nil {
		return err
	}

	return nil
}
