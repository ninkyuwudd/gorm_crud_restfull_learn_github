package router

import (
	"crud/gorm/learn/internal/handler"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Router(db *gorm.DB) {

	router := gin.Default()

	router.POST("/login", handler.LoginHandler(db))
	router.POST("/register", handler.RegisterHandler(db))
	router.GET("/products", handler.ProductsHandler(db))
	router.POST("/products/create", handler.CreateProductDataHandler(db))
	router.DELETE("/products/delete/:id", handler.DeleteProductDataHandler(db))
	router.PUT("/products/update/:id", handler.UpdateProductDataHandler(db))
	router.Run("192.168.1.9:8080")
}
