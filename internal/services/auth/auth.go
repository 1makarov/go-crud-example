package auth

import (
	authorization "github.com/1makarov/go-crud-example/internal/pkg/auth"
)

type Auth struct {
	jwt authorization.JWT
}

func Init(jwt authorization.JWT) *Auth {
	return &Auth{jwt: jwt}
}

func (a *Auth) CreateToken() (string, error) {
	return a.jwt.New()
}

func (a *Auth) ValidToken(token string) error {
	return a.jwt.Valid(token)
}
