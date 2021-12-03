package hash

import (
	"crypto/sha256"
	"fmt"
)

type Manager struct {
	salt string
}

func New(salt string) *Manager {
	return &Manager{salt: salt}
}

func (h *Manager) Hash(password string) (string, error) {
	hash := sha256.New()

	if _, err := hash.Write([]byte(password)); err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", hash.Sum([]byte(h.salt))), nil
}
