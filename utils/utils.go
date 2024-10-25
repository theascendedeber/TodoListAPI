package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password []byte) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	return string(hashedPassword), err
}

func PasswordCheck(password, hashPassword []byte) (bool, error) {
	err := bcrypt.CompareHashAndPassword(hashPassword, password)

	if err != nil {
		return false, err
	}

	return true, nil
}
