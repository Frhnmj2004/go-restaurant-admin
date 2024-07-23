package models

import "github.com/jinzhu/gorm"

type Revenue struct {
	gorm.Model
	TotalRevenue float64 `json:"totalrevenue"`
}
