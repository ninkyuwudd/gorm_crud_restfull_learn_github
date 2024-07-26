package utils

import (
	"crud/gorm/learn/internal/repository"
	"crypto/rand"
	"encoding/hex"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func ComparePasswordHash(db *gorm.DB, email string, pass string) bool {
	getDataUserByEmail, _ := repository.GetDataUserByEmail(db, email)
	err := bcrypt.CompareHashAndPassword([]byte(getDataUserByEmail.Password), []byte(pass))
	return err == nil
}

func GenerateAuthToken() string {
	token, err := randomHex(20)
	if err != nil {
		return "Error When Generate Token"
	}
	return token
}

func randomHex(n int) (string, error) {
	bytes := make([]byte, n)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil

}
