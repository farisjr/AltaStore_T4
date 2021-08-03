package database

import (
	"project/config"
	"project/models"
)

// add new cart
func CreateCart(cart models.Carts) (models.Carts, error) {
	if err := config.DB.Save(&cart).Error; err != nil {
		return cart, err
	}
	return cart, nil
}

// get cart by id
func GetCart(cartId int) (models.Carts, error) {
	var cart models.Carts
	if err := config.DB.Find(&cart, "id=?", cartId).Error; err != nil {
		return cart, err
	}
	return cart, nil
}

// get total price
func GetTotalPrice(cartId int) (int, error) {
	var cartDetails models.CartDetails
	var totalPrice int
	if err := config.DB.Model(&cartDetails).Select("sum(cart_details.price*cart_details.quantity)").Joins("JOIN carts ON carts.id = cart_details.carts_id").Where("carts_id=?", cartId).First(&totalPrice).Error; err == nil {
		return totalPrice, err
	}
	return totalPrice, nil
}

//get total quantity
func GetTotalQty(cartId int) (int, error) {
	var cartDetails models.CartDetails
	var totalQty int
	if err := config.DB.Model(&cartDetails).Select("SUM(cart_details.quantity)").Joins("JOIN carts ON carts.id = cart_details.carts_id").Where("carts_id=?", cartId).First(&totalQty).Error; err == nil {
		return totalQty, err
	}
	return totalQty, nil
}

//update total cart
func UpdateTotalCart(cartId int, newTotalPrice int, newTotalQty int) (models.Carts, error) {
	var cart models.Carts
	if err := config.DB.Model(&cart).Where("id=?", cartId).Updates(models.Carts{TotalPrice: newTotalPrice, TotalQuantity: newTotalQty}).Error; err != nil {
		return cart, err
	}
	return cart, nil
}

//check is cart id exist on table cart
func CheckCartId(cartId int, cart models.Carts) (interface{}, error) {
	if err := config.DB.Where("id=?", cartId).First(&cart).Error; err != nil {
		return nil, err
	}
	return cart.ID, nil
}

// get cart by id
func GetCartById(id int) (models.Carts, error) {
	var cart models.Carts
	if err := config.DB.Find(&cart, "id=?", id).Error; err != nil {
		return cart, err
	}
	return cart, nil
}

//delete cart
func DeleteCart(cartId int) (models.Carts, error) {
	var cart models.Carts
	if err := config.DB.Find(&cart, "id=?", cartId).Unscoped().Delete(&cart).Error; err != nil {
		return cart, err
	}
	return cart, nil
}
