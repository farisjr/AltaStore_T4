package database

import (
	"project/config"
	"project/models"
)

func CreateProductCategories(productCategories models.ProductCategories) (models.ProductCategories, error) {
	if err := config.DB.Save(&productCategories).Error; err != nil {
		return productCategories, err
	}
	return productCategories, nil
}

func GetProductCategories() (interface{}, error) {
	var productcategories []models.ProductCategories
	if err := config.DB.Find(&productcategories).Error; err != nil {
		return nil, err
	}
	return productcategories, nil
}

func GetProductCategoriesId(id int) (models.ProductCategories, error) {
	var productcategories models.ProductCategories
	var count int64
	if err1 := config.DB.Model(&productcategories).Where("id=?", id).Count(&count).Error; count == 0 {
		return productcategories, err1
	}
	if err := config.DB.Find(&productcategories, "id=?", id).Error; err != nil {
		return productcategories, err
	}
	return productcategories, nil
}

func DeleteProductCategoriesById(id int) (interface{}, error) {
	var productcategories models.ProductCategories
	if err := config.DB.Where("id=?", id).Delete(&productcategories).Error; err != nil {
		return nil, err
	}
	return productcategories, nil
}

//update product categories info from database
func UpdateProductCategories(productcategories models.ProductCategories) (models.ProductCategories, error) {
	if tx := config.DB.Save(&productcategories).Error; tx != nil {
		return productcategories, tx
	}
	return productcategories, nil
}

//get 1 specified product categories with productCategories struct output
func GetUpdateProductCategories(id int) models.ProductCategories {
	var productcategories models.ProductCategories
	config.DB.Find(&productcategories, "id=?", id)
	return productcategories
}
