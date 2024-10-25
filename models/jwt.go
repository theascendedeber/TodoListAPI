package models

type JWTHeader struct {
	Alg string `json:"alg"`
	Typ string `json:"typ"`
}

type JWTPayload struct {
	Sub int64 `json:"sub"`
	Exp int64 `json:"exp"`
}
