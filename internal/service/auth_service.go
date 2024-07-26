package service

import (
	"crud/gorm/learn/internal/model"
	"crud/gorm/learn/internal/repository"
	"crud/gorm/learn/utils"
	"log"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var Token []string

func AuthService(db *gorm.DB, email string, pass string) (string, bool) {
	Token = nil
	if utils.ComparePasswordHash(db, email, pass) {

		token := utils.GenerateAuthToken()
		Token = append(Token, token)
		return token, true
	}
	return "", false
}

func RegisterService(db *gorm.DB, userData model.DataUser) error {
	hashPass, err := bcrypt.GenerateFromPassword([]byte(userData.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("Error When Generate Hashing Password", err)
		return err
	}
	dataUserHashPassword := model.DataUser{
		Name:     userData.Name,
		Email:    userData.Email,
		Password: string(hashPass),
	}

	err = repository.RegisterUserData(db, dataUserHashPassword)
	if err != nil {
		log.Println("Error When do register at repository", err)
		return err
	}
	return nil
}
