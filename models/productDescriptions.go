package models

import "gorm.io/gorm"

type ProductDescriptions struct {
	gorm.Model
	Name   string `json:"name" form:"name"`
	Code   string `json:"code" form:"code"`
	Status string `json:"status" form:"status"`
	Price  string `json:"price" form:"price"`
}
