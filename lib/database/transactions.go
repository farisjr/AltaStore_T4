package database

import (
	"project/config"
	"project/models"
)

// get transaction status
func GetTransactionStatus(id int) models.Carts {
	var carts models.Carts
	config.DB.Find(&carts, "id=?", id)
	return carts
}

//udpate transaction status
func UpdateTransactionStatus(cart models.Carts) (models.Carts, error) {

	if err := config.DB.Save(&cart).Error; err != nil {
		return cart, err
	}
	return cart, nil
}
