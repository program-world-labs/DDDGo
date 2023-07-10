package dto

import (
	"crypto/rand"
	"fmt"
)

const idLength = 10

func generateID() (string, error) {
	bytes := make([]byte, idLength)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", bytes), nil
}
