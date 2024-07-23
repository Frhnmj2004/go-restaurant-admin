package models

import "github.com/jinzhu/gorm"

type Order struct {
	gorm.Model
	FoodItemID uint    `json:"fooditemid"`
	Quantity   uint    `json:"quantity"`
	TotalPrice float64 `json:"totalprice"`
}
