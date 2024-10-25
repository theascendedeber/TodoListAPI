package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"time"

	"github.com/theascendedeber/TodoListAPI/models"
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

func GenerateJWT(userID int64) (string, error) {
	secret := "SuperSecret123"
	header := models.JWTHeader{
		Alg: "HS256",
		Typ: "JWT",
	}

	payload := models.JWTPayload{
		Sub: userID,
		Exp: time.Now().Add(time.Minute * 15).Unix(),
	}

	headerJSON, err := json.Marshal(header)
	if err != nil {
		return "", err
	}

	payloadJSON, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}

	encodedHeader := base64.RawURLEncoding.EncodeToString(headerJSON)
	encodedPayload := base64.RawURLEncoding.EncodeToString(payloadJSON)
	signature := createSignature(encodedHeader, encodedPayload, secret)

	return fmt.Sprintf("%s.%s.%s", encodedHeader, encodedPayload, signature), nil
}

func createSignature(header, payload, secret string) string {
	data := fmt.Sprintf("%s.%s", header, payload)
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(data))

	return base64.RawURLEncoding.EncodeToString(h.Sum(nil))
}
