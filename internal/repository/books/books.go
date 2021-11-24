package books

import (
	"context"
	"fmt"
	"github.com/1makarov/go-crud-example/internal/types"
	"github.com/jmoiron/sqlx"
)

type RepoBooks struct {
	db *sqlx.DB
}

func NewRepoBooks(db *sqlx.DB) *RepoBooks {
	return &RepoBooks{db: db}
}

func (r *RepoBooks) Create(ctx context.Context, v types.BookCreateInput) error {
	if _, err := r.db.NamedExecContext(ctx, `

	insert into books (name, title, isbn, description) values (:name, :title, :isbn, :description)

	`, v); err != nil {
		return err
	}

	return nil
}

func (r *RepoBooks) GetByID(ctx context.Context, id int) (*types.Book, error) {
	var v types.Book

	if err := r.db.GetContext(ctx, &v, `select * from books where id = $1`, id); err != nil {
		return nil, err
	}

	return &v, nil
}

func (r *RepoBooks) GetAll(ctx context.Context) ([]types.Book, error) {
	var v []types.Book

	if err := r.db.SelectContext(ctx, &v, `select * from books`); err != nil {
		return nil, err
	}

	if len(v) == 0 {
		v = []types.Book{}
	}

	return v, nil
}

func (r *RepoBooks) DeleteByID(ctx context.Context, id int) error {
	result, err := r.db.ExecContext(ctx, `delete from books where id = $1`, id)
	if err != nil {
		return err
	}

	row, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if row == 0 {
		return fmt.Errorf("book not found")
	}

	return nil
}

func (r *RepoBooks) UpdateByID(ctx context.Context, id int, v types.BookUpdateInput) error {
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

	s := fmt.Sprintf("update books set %s where id = %d", query[:len(query)-1], id)

	result, err := r.db.ExecContext(ctx, s)
	if err != nil {
		return err
	}

	row, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if row == 0 {
		return fmt.Errorf("book not found")
	}

	return nil
}
