package users

import (
	"context"
	"github.com/1makarov/go-crud-example/internal/repository"
	"github.com/1makarov/go-crud-example/internal/types"
	"github.com/1makarov/go-crud-example/pkg/auth"
	"github.com/1makarov/go-crud-example/pkg/hash"
)

type Users struct {
	repo repository.Users
	hash *hash.Manager
	auth *auth.Manager
}

func InitServiceUsers(repo repository.Users, hash *hash.Manager, auth *auth.Manager) *Users {
	return &Users{
		repo: repo,
		hash: hash,
		auth: auth,
	}
}

func (u *Users) SignUp(ctx context.Context, input types.SignUpInput) error {
	passwordHash, err := u.hash.Hash(input.Password)
	if err != nil {
		return err
	}

	input.Password = passwordHash

	return u.repo.SignUp(ctx, input)
}

func (u *Users) SignIn(ctx context.Context, input types.SignInInput) (string, error) {
	passwordHash, err := u.hash.Hash(input.Password)
	if err != nil {
		return "", err
	}

	input.Password = passwordHash

	user, err := u.repo.SignIn(ctx, input)
	if err != nil {
		return "", err
	}

	return u.auth.Create(user.ID)
}

func (u *Users) ParseToken(token string) error {
	return u.auth.Validate(token)
}
