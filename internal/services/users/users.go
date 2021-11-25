package users

import (
	"context"
	"github.com/1makarov/go-crud-example/internal/pkg/hash"
	"github.com/1makarov/go-crud-example/internal/repository"
	"github.com/1makarov/go-crud-example/internal/types"
)

type Users struct {
	repo repository.Users
	hash *hash.Manager
}

func InitServiceUsers(repo repository.Users, hash *hash.Manager) *Users {
	return &Users{
		repo: repo,
		hash: hash,
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

func (u *Users) SignIn(ctx context.Context, input types.SignInInput) (*types.User, error) {
	passwordHash, err := u.hash.Hash(input.Password)
	if err != nil {
		return nil, err
	}

	input.Password = passwordHash

	return u.repo.SignIn(ctx, input)
}
