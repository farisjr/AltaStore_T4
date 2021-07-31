package database

import (
	"project/config"
	"project/models"
)

//function get all sellers table
func GetSellers() (interface{}, error) {
	var sellers []models.Sellers
	if err := config.DB.Find(&sellers).Error; err != nil {
		return nil, err
	}
	return sellers, nil
}

//function to get sellers from id
func GetSellersById(id int) (models.Sellers, error) {
	var sellers models.Sellers
	if err := config.DB.Find(&sellers, "id=?", id).Error; err != nil {
		return sellers, err
	}
	return sellers, nil
}

//function for creating new sellers
func CreateSellers(sellers models.Sellers) (interface{}, error) {
	if err := config.DB.Save(&sellers).Error; err != nil {
		return nil, err
	}
	return sellers, nil
}

//function for deleting sellers by id
func DeleteSellersById(id int) (interface{}, error) {
	var sellers []models.Sellers
	if err := config.DB.Find(&sellers, "id=?", id).Error; err != nil {
		return nil, err
	}
	if err := config.DB.Delete(&sellers, "id=?", id).Error; err != nil {
		return nil, err
	}
	return sellers, nil
}

//function update sellers by id
func UpdateSellersById(sellers models.Sellers, id int) (interface{}, error) {
	if err := config.DB.Find(&sellers, "id=?", id).Error; err != nil {
		return nil, err
	}
	if err := config.DB.Save(&sellers).Error; err != nil {
		return nil, err
	}
	return sellers, nil
}
