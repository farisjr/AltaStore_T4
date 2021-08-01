package database

import (
	"project/config"
	"project/models"
)

//check is product id and cart id exist on table cart_detail
func CheckProductAndCartId(productId, cartId int, cartDetails models.CartDetails) (interface{}, error) {
	if err := config.DB.Where("carts_id=? AND products_id=?", cartId, productId).First(&cartDetails).Error; err != nil {
		return nil, err
	}
	return cartDetails, nil
}

//get cart detail
func GetCartDetailByCartId(cartId int) (models.CartDetails, error) {
	var cartDetails models.CartDetails
	if err := config.DB.Find(&cartDetails, "carts_id=?", cartId).Error; err != nil {
		return cartDetails, err
	}
	return cartDetails, nil
}

//add product to cart
func AddToCart(cartDetails models.CartDetails) (interface{}, error) {
	if err := config.DB.Save(&cartDetails).Error; err != nil {
		return nil, err
	}
	return cartDetails, nil
}

//delete product from cart detail
func DeleteProductFromCart(cartId, productId int) (interface{}, error) {
	var cartDetails models.CartDetails
	if err := config.DB.Find(&cartDetails, "carts_id=? AND products_id=?", cartId, productId).Unscoped().Delete(&cartDetails).Error; err != nil {
		return nil, err
	}
	return cartDetails, nil
}

//get all products from cart detail
func GetListProductCart(cartId int) (interface{}, error) {
	var products []models.Products

	if err := config.DB.Table("products").Preload("Products").Joins("JOIN cart_details ON products.id = cart_details.products_id").Joins("JOIN carts ON cart_details.carts_id = carts.id").Where("carts.id=?", cartId).Find(&products).Error; err != nil {
		return products, nil
	}
	return products, nil
}
