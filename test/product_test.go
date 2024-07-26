package test

import (
	// "crud/gorm/learn/internal/model"
	"crud/gorm/learn/internal/repository"
	"crud/gorm/learn/pkg/database"
	"fmt"

	// "log"
	"testing"

	"github.com/stretchr/testify/assert"
)

var db = database.InitializeDb()

// func TestCreateProduct(t *testing.T) {

// 	product := model.Product{
// 		Name:     "product 07",
// 		Category: "mesin potong",
// 		Quantity: 2,
// 	}

// 	err := repository.CreateNewProductData(db, product)
// 	log.Default()
// 	fmt.Println(err)
// 	assert.NotNil(t, err)

// }

func TestLogin(t *testing.T) {
	// user := model.DataUser{
	// 	Name: "user dev",
	// 	Email: "dev123@gmail.com",
	// 	Password: "dev123",
	// }

	_, err := repository.GetDataUserByEmail(db, "dev123@gmail.com")

	assert.NotNil(t, err)
}

func TestDeleteProduct(t *testing.T) {

	data, err := repository.FindProductById(db, 8)
	fmt.Println(err)
	assert.Nil(t, err)
	assert.Equal(t, 0, len(data))
}
