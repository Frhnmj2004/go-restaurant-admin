package models

import "github.com/jinzhu/gorm"

type Ingredient struct {
	gorm.Model
	FoodItemID uint     `json:"fooditemid"`
	GroceryID  uint     `json:"groceryid"`
	Quantity   float64  `json:"quantity"`
	Grocery    Grocery  `gorm:"foreignkey:GroceryID" json:"-"`
	FoodItem   FoodItem `gorm:"foreignkey:FoodItemID" json:"-"`
}
