package models

import "gorm.io/gorm"

type ProductTypes struct {
	gorm.Model
	ID   int    `gorm:"primaryKey" json:"id" form:"id"`
	Name string `json:"name" form:"name"`
}
