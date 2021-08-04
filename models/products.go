package models

import "gorm.io/gorm"

type Products struct {
	gorm.Model
	ID          int    `gorm:"primaryKey" json:"id" form:"id"`
	Name        string `json:"name" form:"name"`
	Code        string `json:"code" form:"code"`
	Status      string `json:"status" form:"status"`
	Price       int    `json:"price" form:"price"`
	Description string `json:"description" form:"description"`

	//many to many with carts
	Carts []*Carts `gorm:"many2many:cart_details"`

	//1 to many with product category
	ProductCategoriesID int `json:"product_categories_id" form:"product_categories_id"`
}
