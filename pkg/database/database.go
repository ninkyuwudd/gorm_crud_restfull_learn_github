package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitializeDb() *gorm.DB {
	dialect := mysql.Open("root:@tcp(localhost:3306)/gorm_crud_learn?charset=utf8mb4&parseTime=True&loc=Local")
	db, err := gorm.Open(dialect)

	if err != nil {
		panic(fmt.Sprintf("Error %s ,when open db", err))
	}

	return db
}
