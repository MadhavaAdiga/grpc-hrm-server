package utils

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

/*
	Encrypt plain test password with bcrypt algorithm
*/
func HashPassword(password string) (string, error) {
	// uses default cost which is 0
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("failed to hash password: %w", err)
	}

	return string(hashedPassword), nil
}

/*
	Compare a validate the password with encrypted password
*/
func ValidatePassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
