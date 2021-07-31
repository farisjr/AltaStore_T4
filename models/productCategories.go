package models

import "gorm.io/gorm"

type ProductCategories struct {
	gorm.Model
	ID       int        `json:"id" form:"id"`
	Name     string     `json:"name" form:"name"`
	Products []Products `gorm:"foreignKey:ProductCategoriesID"`
}
