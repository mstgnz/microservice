package config

import (
	"golang.org/x/crypto/bcrypt"
)

func HashAndSalt(password []byte) string {
	hash, _ := bcrypt.GenerateFromPassword(password, bcrypt.MinCost)
	return string(hash)
}

func ComparePassword(hashedPassword string, password []byte) bool {
	if bcrypt.CompareHashAndPassword([]byte(hashedPassword), password) != nil {
		return false
	}
	return true
}
