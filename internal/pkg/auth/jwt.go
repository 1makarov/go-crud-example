package auth

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"strconv"
	"time"
)

const errNoSingKey = "empty sign key"

type Manager struct {
	signingKey string
	ttl        time.Duration
}

func New(signingKey string, ttl time.Duration) (*Manager, error) {
	if signingKey == "" {
		return nil, fmt.Errorf(errNoSingKey)
	}

	return &Manager{signingKey: signingKey, ttl: ttl}, nil
}

func (j *Manager) Create(id int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Subject:   strconv.Itoa(id),
		IssuedAt:  time.Now().Unix(),
		ExpiresAt: time.Now().Add(j.ttl).Unix(),
	})

	return token.SignedString([]byte(j.signingKey))
}

func (j *Manager) Validate(token string) error {
	t, err := jwt.Parse(token, func(token *jwt.Token) (i interface{}, err error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(j.signingKey), nil
	})

	if err != nil {
		return err
	}

	return t.Claims.Valid()
}
