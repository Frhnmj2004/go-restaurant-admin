package models

import (
	"github.com/jinzhu/gorm"
)

type FoodItem struct {
	gorm.Model
	Name        string       `gorm:"unique" json:"name"`
	Price       float64      `json:"price"`
	Ingredients []Ingredient `gorm:"many2many:fooditem_ingredients" json:"ingredients"`
}
