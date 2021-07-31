package database

import (
	"project/config"
	"project/models"
)

//function get all payment methods table
func GetPaymentMethods() (interface{}, error) {
	var paymentMethods []models.PaymentMethods
	if err := config.DB.Find(&paymentMethods).Error; err != nil {
		return nil, err
	}
	return paymentMethods, nil
}

//function to get payment methods from id
func GetPaymentMethodsById(id int) (models.PaymentMethods, error) {
	var paymentMethods models.PaymentMethods
	if err := config.DB.Find(&paymentMethods, "id=?", id).Error; err != nil {
		return paymentMethods, err
	}
	return paymentMethods, nil
}

//function for creating new payment methods
func CreatePaymentMethods(createPaymentMethod models.PaymentMethods) (interface{}, error) {
	if err := config.DB.Save(&createPaymentMethod).Error; err != nil {
		return nil, err
	}
	return createPaymentMethod, nil
}

//function for deleting payment methods by id
func DeletePaymentMethodsById(id int) (interface{}, error) {
	var paymentMethods []models.PaymentMethods
	if err := config.DB.Find(&paymentMethods, "id=?", id).Error; err != nil {
		return nil, err
	}
	if err := config.DB.Delete(&paymentMethods, "id=?", id).Error; err != nil {
		return nil, err
	}
	return paymentMethods, nil
}

//function update payment methods by id
func UpdatePaymentMethodsById(paymentMethods models.PaymentMethods, id int) (interface{}, error) {
	if err := config.DB.Find(&paymentMethods, "id=?", id).Error; err != nil {
		return nil, err
	}
	if err := config.DB.Save(&paymentMethods).Error; err != nil {
		return nil, err
	}
	return paymentMethods, nil
}
