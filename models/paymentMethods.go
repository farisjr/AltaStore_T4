package models

import "gorm.io/gorm"

type PaymentMethods struct {
	gorm.Model
	Name   string `json:"name" form:"name"`
	Status string `json:"status" form:"status"`
}
