package models

import "gorm.io/gorm"

type Customers struct {
	gorm.Model
	Name        string `json:"name" form:"name"`
	Address     string `json:"address" form:"address"`
	DateOfBirth string `json:"date_of_birth" form:"date_of_birth"`
	Status      string `json:"status" form:"status"`
	Gender      string `json:"gender" form:"gender"`
}
