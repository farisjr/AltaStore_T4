package database

import (
	"project/config"
	"project/models"
)

//check is product exist on table product
func CheckProductId(productId int, product models.Products) (interface{}, error) {
	if err := config.DB.Where("id=?", productId).First(&product).Error; err != nil {
		return nil, err
	}
	return product.ID, nil
}

// get product by id
func GetProduct(productId int) (models.Products, error) {
	var product models.Products
	if err := config.DB.Find(&product, "id=?", productId).Error; err != nil {
		return product, err
	}
	return product, nil
}
