package main

import (
	"crud/gorm/learn/internal/router"
	"crud/gorm/learn/pkg/database"
)

func main() {

	db := database.InitializeDb()

	router.Router(db)
}
