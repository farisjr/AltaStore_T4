package models

import "gorm.io/gorm"

type Transactions struct {
	gorm.Model
	Status string `json:"status" form:"status"`
}
