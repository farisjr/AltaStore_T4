package models

import "gorm.io/gorm"

type ProductTypes struct {
	gorm.Model
	Name string `json:"name" form:"name"`
}
