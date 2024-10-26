package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password []byte) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	return string(hashedPassword), err
}

func VerifyPassword(password, hashPassword []byte) bool {
	err := bcrypt.CompareHashAndPassword(hashPassword, password)

	if err != nil {
		return false
	}

	return true
}
