package controllers

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type FoodItemController struct {
	DB *gorm.DB
}

func NewFoodItemController(db *gorm.DB) *FoodItemController {
	return &FoodItemController{DB: db}
}

func (ctrl *FoodItemController) CreateFoodItem(ctx *fiber.Ctx) error {

	return nil
}

func (ctrl *FoodItemController) GetFoodItemByID(ctx *fiber.Ctx) error {

	return nil
}

func (ctrl *FoodItemController) GetAllFoodItems(ctx *fiber.Ctx) error {

	return nil
}
