package database

import (
	"project/config"
	"project/models"
)

func CheckProductId(productId int, product models.Products) (interface{}, error) {
	if err := config.DB.Where("id=?", productId).First(&product).Error; err != nil {
		return nil, err
	}
	return product.ID, nil
}

func GetProduct(productId int) (models.Products, error) {
	var product models.Products
	if err := config.DB.Find(&product, "id=?", productId).Error; err != nil {
		return product, err
	}
	return product, nil
}

func GetCart(cartId int) (models.Carts, error) {
	var cart models.Carts
	if err := config.DB.Find(&cart, "id=?", cartId).Error; err != nil {
		return cart, err
	}
	return cart, nil
}

func AddToCart(cartDetails models.CartDetails) (interface{}, error) {
	if err := config.DB.Save(&cartDetails).Error; err != nil {
		return nil, err
	}
	return cartDetails, nil
}
