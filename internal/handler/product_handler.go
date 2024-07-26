package handler

import (
	"crud/gorm/learn/internal/model"
	"crud/gorm/learn/internal/service"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ProductsHandler(db *gorm.DB) gin.HandlerFunc {

	return func(ctx *gin.Context) {
		bearerToken := ctx.Request.Header.Get("Authorization")
		getToken := strings.Split(bearerToken, " ")[1]
		tokenValid := false

		for _, token := range service.Token {
			if token == getToken {
				tokenValid = true
				dataProducts, err := service.GetProductsDataService(db)

				if err != nil {
					ctx.JSON(http.StatusUnauthorized, gin.H{
						"message": "Error when get data product",
					})
					return
				}

				ctx.IndentedJSON(http.StatusOK, gin.H{
					"data": dataProducts,
				
					
				})

				return
			}
		}

		if !tokenValid {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized",
			})
		}

	}
}

func CreateProductDataHandler(db *gorm.DB) gin.HandlerFunc {

	return func(ctx *gin.Context) {

		bearerToken := ctx.Request.Header.Get("Authorization")
		getToken := strings.Split(bearerToken, " ")[1]

		var dataProduct model.Product

		if err := ctx.ShouldBindJSON(&dataProduct); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": fmt.Sprintf("error input data : %s", err),
			})
			return
		}

		for _, token := range service.Token {
			if token == getToken {
				err := service.CreateProductData(db, dataProduct)

				if err != nil {
					ctx.JSON(http.StatusInternalServerError, gin.H{
						"message": "Error when Create New Data",
					})
				}

				ctx.JSON(http.StatusOK, gin.H{
					"message": "Create New Product Data Successfully",
				})
				return
			} else {
				ctx.JSON(http.StatusUnauthorized, gin.H{
					"message": "Unauthorized",
				})
				return
			}
		}

	}

}

func DeleteProductDataHandler(db *gorm.DB) gin.HandlerFunc {

	return func(ctx *gin.Context) {
		bearerToken := ctx.Request.Header.Get("Authorization")
		getToken := strings.Split(bearerToken, " ")[1]

		idProduct, _ := strconv.Atoi(ctx.Param("id"))

		if err := ctx.ShouldBindJSON(&idProduct); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": fmt.Sprintf("error input data : %s", err),
			})
			return
		}

		for _, token := range service.Token {
			if token == getToken {
				dataLen, err := service.DeleteProductData(db, idProduct)
				if dataLen != 0 && err == nil {
					ctx.JSON(http.StatusOK, gin.H{
						"message": "Delete Product Data Successfully",
					})
				} else if dataLen == 0 && err == nil {
					ctx.JSON(http.StatusOK, gin.H{
						"message": "Id product not found",
					})
				}

				if err != nil {
					ctx.JSON(http.StatusInternalServerError, gin.H{
						"message": "Data not foud or Error when Delete Data ",
					})
				}
				return
			} else {
				ctx.JSON(http.StatusUnauthorized, gin.H{
					"message": "Unauthorized",
				})
			}
		}

	}
}

func UpdateProductDataHandler(db *gorm.DB) gin.HandlerFunc {

	return func(ctx *gin.Context) {
		bearerToken := ctx.Request.Header.Get("Authorization")
		getToken := strings.Split(bearerToken, " ")[1]

		idProduct, _ := strconv.Atoi(ctx.Param("id"))

		var productData model.ProductUpdate

		if err := ctx.ShouldBindJSON(&productData); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": fmt.Sprintf("error when input data: %s", err),
			})
		}

		for _, token := range service.Token {
			if token == getToken {
				dataLen, err := service.UpdateProductData(db, idProduct, productData)

				if dataLen == 0 && err == nil {
					ctx.JSON(http.StatusNotFound, gin.H{
						"message": "Cannot find Id data",
					})
				} else if dataLen != 0 {
					ctx.JSON(http.StatusOK, gin.H{
						"message": "Update Data Successfully",
					})
				}

				if err != nil {
					ctx.JSON(http.StatusOK, gin.H{
						"message": fmt.Sprintf("Error When Update Data : %s", err),
					})
				}
			}else {
				ctx.JSON(http.StatusUnauthorized,gin.H{
					"message" : "Unauthorized",
				})
			}
		}
	}
}
