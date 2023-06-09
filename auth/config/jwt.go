package config

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// GetSecretKey get key
func GetSecretKey() string {
	secretKey := os.Getenv("JWT_SECRET")
	if secretKey != "" {
		secretKey = "my_secret"
	}
	return secretKey
}

// GenerateToken token generate
func GenerateToken(userId uint) (string, error) {
	claims := jwt.RegisteredClaims{
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		ExpiresAt: jwt.NewNumericDate(time.Now().AddDate(0, 0, 1)),
		Issuer:    strconv.Itoa(int(userId)),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(GetSecretKey()))
	if err != nil {
		return "", err
	}
	return t, nil
}

// ValidateToken token validate
func ValidateToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New(fmt.Sprintf("unexpected signing method %v", t.Header["alg"]))
		}
		return []byte(GetSecretKey()), nil
	})
}

func GetUserIDByToken(token string) (uint, error) {
	aToken, err := ValidateToken(token)
	if err != nil {
		return 0, err
	}
	claims := aToken.Claims.(jwt.MapClaims)
	id, _ := strconv.ParseUint(fmt.Sprintf("%v", claims["iss"]), 10, 32)
	return uint(id), nil
}
