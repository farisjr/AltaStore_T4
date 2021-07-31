package models

import "gorm.io/gorm"

type Carts struct {
	gorm.Model
	ID                 int    `json:"id" form:"id"`
	StatusTransactions string `json:"status_transactions" form:"status_transactions"`
	TotalQuantity      int    `json:"total_quantity" form:"total_quantity"`
	TotalPrice         int    `json:"total_price" form:"total_price"`

	//many to many
	Products []*Products `gorm:"many2many:cart_details"`

	//1 to many
	CustomersID      int
	PaymentMethodsID int
}
