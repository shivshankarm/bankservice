package util

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hashedpassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("error in generating bcrypt hashpassword %w", err)
	}
	return string(hashedpassword), nil
}

func CheckPassword(password string, hashedpassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedpassword), []byte(password))
}
