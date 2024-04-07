package middleware

import (
	"golang.org/x/crypto/bcrypt"
)

func ComparePassword(hashedPassword []byte, inputPassword string) error {
	return bcrypt.CompareHashAndPassword(hashedPassword, []byte(inputPassword))

}
