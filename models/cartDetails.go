package models

import "gorm.io/gorm"

type CartDetails struct {
	gorm.Model
	Subtotal string `json:"subtotal" form:"subtotal"`
}
