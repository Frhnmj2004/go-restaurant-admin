package models

import "github.com/jinzhu/gorm"

type Grocery struct {
	gorm.Model
	Name     string  `gorm:"unique" json:"name"`
	Quantity float64 `json:"quantity"`
}
