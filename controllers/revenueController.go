package controllers

import (
	"net/http"

	"github.com/Frhnmj2004/restaurant-admin/helper"
	"github.com/Frhnmj2004/restaurant-admin/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type RevenueController struct {
	DB *gorm.DB
}

func NewRevenueController(db *gorm.DB) *RevenueController {
	return &RevenueController{DB: db}
}

func (ctrl *RevenueController) GetRevenue(ctx *fiber.Ctx) error {
	revenue := &models.Revenue{}

	err := ctrl.DB.First(revenue).Error
	if err != nil {
		return helper.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to get revenue")
	}

	ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "Revenue retrieved successfully",
		"data":    revenue,
	})

	return nil
}
