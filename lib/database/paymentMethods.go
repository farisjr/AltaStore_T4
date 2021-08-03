package database

import (
	"project/config"
	"project/models"
)

//check is payment exist on table payment methods
func CheckPayment(paymentId int, payment models.PaymentMethods) (interface{}, error) {
	if err := config.DB.Where("id=?", paymentId).First(&payment).Error; err != nil {
		return nil, err
	}
	return payment.ID, nil
}

//function get all payment methods table
func GetPaymentMethods() (interface{}, error) {
	var paymentMethods []models.PaymentMethods
	if err := config.DB.Find(&paymentMethods).Error; err != nil {
		return nil, err
	}
	return paymentMethods, nil
}

// get payment methods from id
func GetPaymentMethodsById(id int) (models.PaymentMethods, error) {
	var paymentMethods models.PaymentMethods
	if err := config.DB.Find(&paymentMethods, "ID=?", id).Error; err != nil {
		return paymentMethods, err
	}
	return paymentMethods, nil
}

// creating new payment methods
func CreatePaymentMethods(createPaymentMethod models.PaymentMethods) (models.PaymentMethods, error) {
	if err := config.DB.Save(&createPaymentMethod).Error; err != nil {
		return createPaymentMethod, err
	}
	return createPaymentMethod, nil
}

// deleting payment methods by id
func DeletePaymentMethodsById(id int) (interface{}, error) {
	var paymentMethods []models.PaymentMethods
	if err := config.DB.Find(&paymentMethods, "ID=?", id).Error; err != nil {
		return nil, err
	}
	if err := config.DB.Delete(&paymentMethods, "ID=?", id).Error; err != nil {
		return nil, err
	}
	return paymentMethods, nil
}

//get 1 specified payment method
func GetUpdatePaymentMethod(id int) models.PaymentMethods {
	var paymentMethod models.PaymentMethods
	config.DB.Find(&paymentMethod, "id=?", id)
	return paymentMethod
}

//update payment method  from database
func UpdatePaymentMethod(paymentMethod models.PaymentMethods) (models.PaymentMethods, error) {
	if tx := config.DB.Save(&paymentMethod).Error; tx != nil {
		return paymentMethod, tx
	}
	return paymentMethod, nil
}
