package controllers

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type OrderController struct {
	DB *gorm.DB
}

func NewOrderController(db *gorm.DB) *OrderController {
	return &OrderController{DB: db}
}

func (ctrl *OrderController) GetAllOrders(c *fiber.Ctx) error {

	return nil
}

func (ctrl *OrderController) GetOrderByID(c *fiber.Ctx) error {

	return nil
}

func (ctrl *OrderController) PlaceOrder(c *fiber.Ctx) error {

	return nil
}
