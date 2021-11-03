package authorization

import (
	"fmt"
	"github.com/1makarov/go-crud-example/internal/pkg/auth/jwt"
	"time"
)

const errNoSalt = "empty salt"

type JWT interface {
	New() (string, error)
	Valid(token string) error
	Refresh() (string, error)
}

type Auth struct {
	JWT
}

func New(salt string, ttl time.Duration) (*Auth, error) {
	if salt == "" {
		return nil, fmt.Errorf(errNoSalt)
	}

	return &Auth{
		JWT: jwt.Init(salt, ttl),
	}, nil
}
