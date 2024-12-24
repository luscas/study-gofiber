package jwt

import (
	"fmt"
	"log"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

var secretKey []byte

func InitSecretKey() {
	secretKey = []byte(os.Getenv("JWT_SECRET"))

	if len(secretKey) == 0 {
		log.Fatal("JWT_SECRET environment variable not found")
	}
}

type Payload struct {
	Sub  string `json:"sub"`
	Name string `json:"name"`
}

func GenerateToken(payload Payload) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":  payload.Sub,
		"name": payload.Name,
	})

	return token.SignedString(secretKey)
}

func VerifyToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return err
	}

	if !token.Valid {
		return fmt.Errorf("invalid token")
	}

	return nil
}
