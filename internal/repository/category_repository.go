package repository

import (
	"crud/gorm/learn/internal/model"
	"log"

	"gorm.io/gorm"
)

func GetCategoryData(db *gorm.DB, id int) (model.Category, error) {

	var dataCategory model.Category

	err := db.Take(&dataCategory, "id =?", id).Error

	if err != nil {
		log.Printf("data null /error")
		return dataCategory, err
	}
	log.Printf(dataCategory.Name)

	return dataCategory, nil
}
