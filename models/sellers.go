package models

import "gorm.io/gorm"

type Sellers struct {
	gorm.Model
	Name string `json:"name" form:"name"`
}
