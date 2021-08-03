package database

import (
	"project/config"
	"project/middlewares"
	"project/models"
)

//get all seller
func GetAllSellers() (interface{}, error) {
	var sellers []models.Sellers
	if err := config.DB.Find(&sellers).Error; err != nil {
		return nil, err
	}
	return sellers, nil
}

// get one seller
func GetOneSeller(id int) (interface{}, error) {
	var seller models.Sellers
	var count int64
	if err1 := config.DB.Model(&seller).Where("id=?", id).Count(&count).Error; count == 0 {
		return nil, err1
	}
	if err := config.DB.Find(&seller, "id=?", id).Error; err != nil {
		return nil, err
	}
	return seller, nil
}

//create new seller
func CreateSeller(sellers models.Sellers) (interface{}, error) {
	if err := config.DB.Save(&sellers).Error; err != nil {
		return nil, err
	}
	return sellers, nil
}

//login seller with matching data from database
func LoginSellers(email, password string) (interface{}, error) {
	var seller models.Sellers
	var err error
	if err = config.DB.Where("email = ? AND password = ?", email, password).First(&seller).Error; err != nil {
		return nil, err
	}
	seller.Token, err = middlewares.CreateToken(int(seller.ID))
	if err != nil {
		return nil, err
	}
	if err := config.DB.Save(seller).Error; err != nil {
		return nil, err
	}
	return seller, err
}

//delete seller by id
func DeleteSellers(id int) (interface{}, error) {
	var sellers models.Sellers
	if err := config.DB.Where("id=?", id).Delete(&sellers).Error; err != nil {
		return nil, err
	}
	return sellers, nil
}

//update seller info from database
func UpdateSellers(sellers models.Sellers) (interface{}, error) {
	if tx := config.DB.Save(&sellers).Error; tx != nil {
		return nil, tx
	}
	return sellers, nil
}

//get 1 specified user with User struct output
func GetUpdateSellers(id int) models.Sellers {
	var sellers models.Sellers
	config.DB.Find(&sellers, "id=?", id)
	return sellers
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
