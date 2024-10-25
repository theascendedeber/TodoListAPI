package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/theascendedeber/TodoListAPI/models"
)

const secret = "SuperSecret123"

func GenerateJWT(userID int64) (string, error) {

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

func VerifyJWT(token string) (bool, error) {
	parts := splitToken(token)

	if len(parts) != 3 {
		return false, fmt.Errorf("Недействительный токен")
	}

	header := parts[0]
	payload := parts[1]
	expectedSignature := createSignature(header, payload, secret)

	if expectedSignature == parts[2] {
		return true, nil
	}
	return false, nil
}

func createSignature(header, payload, secret string) string {
	data := fmt.Sprintf("%s.%s", header, payload)
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(data))

	return base64.RawURLEncoding.EncodeToString(h.Sum(nil))
}

func splitToken(token string) []string {
	return strings.Split(token, ".")
}
