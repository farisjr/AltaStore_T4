package models

import "gorm.io/gorm"

type Carts struct {
	gorm.Model
	Status        string `json:"status" form:"status"`
	TotalQuantity string `json:"total_quantity" form:"total_quantity"`
	TotalPrice    string `json:"total_price" form:"total_price"`
}
