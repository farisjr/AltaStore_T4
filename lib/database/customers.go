package database

import (
	"project/config"
	"project/middlewares"
	"project/models"
)

func GetCustomers(id int) (interface{}, error) {
	var customer models.Customers
	var count int64
	if err1 := config.DB.Model(&customer).Where("id=?", id).Count(&count).Error; count == 0 {
		return nil, err1
	}
	if err := config.DB.Find(&customer, "id=?", id).Error; err != nil {
		return nil, err
	}
	return customer, nil
}

func CreateCustomer(customers models.Customers) (interface{}, error) {
	if err := config.DB.Save(&customers).Error; err != nil {
		return nil, err
	}
	return customers, nil
}

//login customer with matching data from database
func LoginCustomers(email, password string) (interface{}, error) {
	var customer models.Customers
	var err error
	if err = config.DB.Where("email = ? AND password = ?", email, password).First(&customer).Error; err != nil {
		return nil, err
	}
	customer.Token, err = middlewares.CreateToken(int(customer.ID))
	if err != nil {
		return nil, err
	}
	if err := config.DB.Save(customer).Error; err != nil {
		return nil, err
	}
	return customer, err
}
