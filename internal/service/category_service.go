package service

import (
	"crud/gorm/learn/internal/model"
	"crud/gorm/learn/internal/repository"

	"gorm.io/gorm"
)

func GetCategoryDataService(db *gorm.DB, id int) (model.Category,error) {
	data,err := repository.GetCategoryData(db, id)

	return data,err
}
