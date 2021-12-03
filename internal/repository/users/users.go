package users

import (
	"context"
	"github.com/1makarov/go-crud-example/internal/types"
	"github.com/jmoiron/sqlx"
)

type RepoUsers struct {
	db *sqlx.DB
}

func NewRepoUsers(db *sqlx.DB) *RepoUsers {
	return &RepoUsers{db: db}
}

func (r *RepoUsers) SignUp(ctx context.Context, input types.SignUpInput) error {
	_, err := r.db.ExecContext(ctx, `

	insert into users (name, email, password_hash) values ($1, $2, $3)
	
	`, input.Name, input.Email, input.Password)

	return err
}

func (r *RepoUsers) SignIn(ctx context.Context, input types.SignInInput) (types.User, error) {
	var user types.User

	err := r.db.GetContext(ctx, &user, `

	select * from users where email = $1 and password_hash = $2
	
	`, input.Email, input.Password)

	return user, err
}
