package dto

import (
	"crypto/rand"
	"fmt"
	"time"
)

const idLength = 10

func generateID() (string, error) {
	bytes := make([]byte, idLength)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", bytes), nil
}

func ParseDateString(data map[string]interface{}) error {
	if tm, ok := data["created_at"].(string); ok {
		t, err := time.Parse(time.RFC3339Nano, tm)
		if err != nil {
			return err
		}

		data["created_at"] = t
	}

	if tm, ok := data["updated_at"].(string); ok {
		t, err := time.Parse(time.RFC3339Nano, tm)
		if err != nil {
			return err
		}

		data["updated_at"] = t
	}

	if tm, ok := data["deleted_at"].(string); ok {
		t, err := time.Parse(time.RFC3339Nano, tm)
		if err != nil {
			return err
		}

		data["deleted_at"] = t
	}

	return nil
}
