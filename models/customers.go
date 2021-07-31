package models

import "gorm.io/gorm"

type Customers struct {
	gorm.Model
	ID       int    `json:"id" form:"id"`
	Name     string `json:"name" form:"name"`
	Address  string `json:"address" form:"address"`
	Gender   string `json:"gender" form:"gender"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Token    string `json:"token" form:"token"`

	//1 to many with carts
	Carts []Carts `gorm:"foreignKey:CustomersID"`
}
