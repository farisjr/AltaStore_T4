package database

import (
	"project/config"
	"project/models"
)

func CreateCart(cart models.Carts) (interface{}, error) {
	if err := config.DB.Save(&cart).Error; err != nil {
		return nil, err
	}
	return cart, nil
}

func GetTotalPrice(cartId int) (int, error) {
	var cartDetails models.CartDetails
	var totalPrice int
	if err := config.DB.Model(&cartDetails).Select("sum(cart_details.price*cart_details.quantity)").Joins("JOIN carts ON carts.id = cart_details.carts_id").Where("carts_id=?", cartId).First(&totalPrice).Error; err == nil {
		return totalPrice, err
	}
	return totalPrice, nil
}

func GetTotalQty(cartId int) (int, error) {
	var cartDetails models.CartDetails
	var totalQty int
	if err := config.DB.Model(&cartDetails).Select("SUM(cart_details.quantity)").Joins("JOIN carts ON carts.id = cart_details.carts_id").Where("carts_id=?", cartId).First(&totalQty).Error; err == nil {
		return totalQty, err
	}
	return totalQty, nil
}

func UpdateTotalCart(cartId int, newTotalPrice int, newTotalQty int) (interface{}, error) {
	var cart models.Carts
	if err := config.DB.Model(&cart).Where("id=?", cartId).Updates(models.Carts{TotalPrice: newTotalPrice, TotalQuantity: newTotalQty}).Error; err != nil {
		return nil, err
	}
	return cart, nil
}

func CheckCartId(cartId int, cart models.Carts) (interface{}, error) {
	if err := config.DB.Where("id=?", cartId).First(&cart).Error; err != nil {
		return nil, err
	}
	return cart.ID, nil
}

func GetCartById(id int) (interface{}, error) {
	var cart models.Carts
	if err := config.DB.Find(&cart, "id=?", id).Error; err != nil {
		return nil, err
	}
	return cart, nil
}

func GetListProductCart(cartId int) (interface{}, error) {
	var products []models.Products

	if err := config.DB.Table("products").Preload("Products").Joins("JOIN cart_details ON products.id = cart_details.products_id").Joins("JOIN carts ON cart_details.carts_id = carts.id").Where("carts.id=?", cartId).Find(&products).Error; err != nil {
		return products, nil
	}
	return products, nil
}
