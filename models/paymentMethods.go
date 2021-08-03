package models

import "gorm.io/gorm"

type PaymentMethods struct {
	gorm.Model
	ID     int    `gorm:"primaryKey"`
	Name   string `json:"name" form:"name"`
	Status string `json:"status" form:"status"`

	//1 to many with carts
	Carts []Carts `gorm:"foreignKey:PaymentMethodsID"`
}
