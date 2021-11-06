package auth

import (
	"fmt"
	"github.com/1makarov/go-crud-example/internal/pkg/auth/jwt"
	"time"
)

const errNoSingKey = "empty sign key"

type JWT interface {
	New() (string, error)
	Valid(token string) error
}

type Auth struct {
	JWT
}

func New(salt string, ttl time.Duration) (*Auth, error) {
	if salt == "" {
		return nil, fmt.Errorf(errNoSingKey)
	}

	return &Auth{
		JWT: jwt.Init(salt, ttl),
	}, nil
}
