package models

import "gorm.io/gorm"

type ProductDescriptions struct {
	gorm.Model
	ID     int    `gorm:"primaryKey" json:"id" form:"id"`
	Name   string `json:"name" form:"name"`
	Code   string `json:"code" form:"code"`
	Status string `json:"status" form:"status"`
	Price  string `json:"price" form:"price"`
}
