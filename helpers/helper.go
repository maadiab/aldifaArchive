package helpers

import (
	"golang.org/x/crypto/bcrypt"
)

func ServeTemplates(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

// Serve HTML Templates
