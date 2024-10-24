package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"time"

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

func GenerateJWT(id_user int, name_user string) string {
	secret := "some_secret123"

	header := "{\"alg\": \"HS256\", \"typ\": \"JWT\"}"
	encodeHeader := base64.RawURLEncoding.EncodeToString([]byte(header))

	iat := time.Now().Unix()
	payload := fmt.Sprintf("{\"sub\":%d, \"name\":%s, \"iat\":%d}", id_user, name_user, iat)
	encodedPayload := base64.RawURLEncoding.EncodeToString([]byte(payload))

	data := fmt.Sprintf("%s.%s", encodeHeader, encodedPayload)
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(data))
	signature := base64.RawURLEncoding.EncodeToString(h.Sum(nil))

	token := fmt.Sprintf("%s.%s.%s", encodeHeader, encodedPayload, signature)
	return token
}
