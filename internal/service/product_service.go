package service

import (
	"crud/gorm/learn/internal/model"
	"crud/gorm/learn/internal/repository"

	"gorm.io/gorm"
)

func GetProductsDataService(db *gorm.DB)( []model.Product,error) {
	dataProducts, err := repository.GetProductsData(db)

	if err != nil {
		return nil,err
	}

	return dataProducts,nil
}

func CreateProductData(db *gorm.DB, dataProduct model.Product) error {

	err := repository.CreateNewProductData(db, dataProduct)

	return err
}

func DeleteProductData(db *gorm.DB, dataProduct int) (int, error) {

	data, _ := repository.FindProductById(db, dataProduct)

	if len(data) != 0 {
		err := repository.DeleteProductData(db, dataProduct)

		if err != nil {
			return 0, err
		}
		return len(data), err
	}

	return 0, nil
}

func UpdateProductData(db *gorm.DB, id int,dataUpdate model.ProductUpdate) (int, error) {
	data, _ := repository.FindProductById(db, id)

	if len(data) != 0 {
		
		err := repository.UpdateProductById(db, data[0], dataUpdate)

		if err != nil {
			return 0, err
		}

		return len(data), err
	}

	return 0, nil
}
