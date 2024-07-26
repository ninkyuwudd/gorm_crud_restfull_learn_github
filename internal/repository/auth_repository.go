package repository

import (
	"crud/gorm/learn/internal/model"

	"fmt"

	"gorm.io/gorm"
)

func GetDataUserByEmail(db *gorm.DB, email string) (model.DataUser, error) {
	var dataUser model.DataUser
	err := db.Take(&dataUser, "email = ?", email).Error
	if err != nil {
		return dataUser, fmt.Errorf("failed to find users: %w", err)
	}

	return dataUser, err
}

func RegisterUserData(db *gorm.DB, userModel model.DataUser) error {

	err := db.Create(&userModel).Error
	if err != nil {
		return fmt.Errorf("failed to find users: %w", err)
	}

	return err
}
