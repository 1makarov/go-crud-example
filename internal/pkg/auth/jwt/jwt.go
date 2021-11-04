package jwt

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type JWT struct {
	signingKey string
	ttl        time.Duration
}

func Init(signingKey string, ttl time.Duration) *JWT {
	return &JWT{signingKey: signingKey, ttl: ttl}
}

func (j *JWT) New() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: time.Now().Add(j.ttl).Unix(),
	})

	return token.SignedString([]byte(j.signingKey))
}

func (j *JWT) Valid(token string) error {
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
