package models

import "time"

type CartDetails struct {
	ProductsID int `gorm:"primaryKey" json:"products_id" form:"products_id"`
	CartsID    int `gorm:"primaryKey" json:"carts_id" form:"carts_id"`
	Quantity   int `json:"quantity" form:"quantity"`
	Price      int `json:"price" form:"price"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
