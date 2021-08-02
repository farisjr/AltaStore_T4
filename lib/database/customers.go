package database

import (
	"project/config"
	"project/middlewares"
	"project/models"
)

func GetCustomersid(id int) (interface{}, error) {
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

func DeleteCustomersById(id int) (interface{}, error) {
	var customers models.Customers
	if err := config.DB.Where("id=?", id).Delete(&customers).Error; err != nil {
		return nil, err
	}
	return customers, nil
}

//update user info from database
func UpdateCustomers(customers models.Customers) (interface{}, error) {
	if tx := config.DB.Save(&customers).Error; tx != nil {
		return nil, tx
	}
	return customers, nil
}

//get 1 specified user with User struct output
func GetUpdateCustomers(id int) models.Customers {
	var customers models.Customers
	config.DB.Find(&customers, "id=?", id)
	return customers
}
