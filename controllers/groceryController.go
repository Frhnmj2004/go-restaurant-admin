package controllers

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type GroceryController struct {
	DB *gorm.DB
}

func NewGroceryController(db *gorm.DB) *GroceryController {
	return &GroceryController{DB: db}
}

func (ctrl *GroceryController) CreateGrocery(ctx *fiber.Ctx) error {

	return nil
}

func (ctrl *GroceryController) GetAllGroceries(ctx *fiber.Ctx) error {

	return nil
}

func (ctrl *GroceryController) GetGroceryByID(ctx *fiber.Ctx) error {

	return nil
}

func (ctrl *GroceryController) UpdateGrocery(ctx *fiber.Ctx) error {

	return nil
}
