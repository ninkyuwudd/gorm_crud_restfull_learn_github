package repository

import (
	"crud/gorm/learn/internal/model"
	"fmt"
	"log"
	"time"

	"gorm.io/gorm"
)

func GetProductsData(db *gorm.DB) ([]model.Product, error) {
	var productSData []model.Product

	// err := db.Find(&productSData).Error
	err := db.Joins("Category").Find(&productSData).Error

	if err != nil {
		log.Println("Error When select * products")
		return nil, err
	}

	return productSData, nil
}

func CreateNewProductData(db *gorm.DB, dataProduct model.Product) error {

	err := db.Create(&dataProduct).Error

	if err != nil {
		return fmt.Errorf("failed to create product: %w", err)
	}

	return nil
}

func DeleteProductData(db *gorm.DB, id int) error {

	err := db.Delete(&model.Product{}, id).Error

	if err != nil {
		return fmt.Errorf("failed to delete: %w", err)
	}

	return nil
}

func FindProductById(db *gorm.DB, id int) ([]model.Product, error) {
	var product []model.Product
	if err := db.Find(&product, id).Error; err != nil {
		return nil, err
	}
	fmt.Println(len(product))
	return product, nil
}

func UpdateProductById(db *gorm.DB, dataProduct model.Product, dataUpdate model.ProductUpdate) error {

	err := db.Model(&dataProduct).Updates(model.ProductUpdate{
		Name:      dataUpdate.Name,
		Category:  dataUpdate.Category,
		Quantity:  dataUpdate.Quantity,
		UpdatedAt: time.Now(),
	}).Error

	if err != nil {
		return fmt.Errorf("failed to update: %w", err)
	}

	return nil
}
