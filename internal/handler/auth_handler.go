package handler

import (
	"crud/gorm/learn/internal/model"
	"crud/gorm/learn/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func LoginHandler(db *gorm.DB) gin.HandlerFunc {

	return func(ctx *gin.Context) {
		var dataUser model.DataUser

		if err := ctx.ShouldBindJSON(&dataUser); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		email := dataUser.Email
		pass := dataUser.Password

		token, isLogin := service.AuthService(db, email, pass)

		if isLogin {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "Login Successfully",
				"token":   token,
			})
		} else {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "Bad request",
			})
		}

	}

}

func RegisterHandler(db *gorm.DB) gin.HandlerFunc {

	return func(ctx *gin.Context) {
		var dataUser model.DataUser

		if err := ctx.ShouldBindJSON(&dataUser); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err := service.RegisterService(db,dataUser)

		if(err != nil){
			ctx.JSON(http.StatusBadRequest,gin.H{
				"error" : err.Error(),
			})
		}else{
			ctx.JSON(http.StatusOK,gin.H{
				"message" : "User Register Successfully",
			})
		}

	}

}
