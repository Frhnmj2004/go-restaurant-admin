package controllers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"github.com/Frhnmj2004/restaurant-admin/helper"
	"github.com/Frhnmj2004/restaurant-admin/models"
	"github.com/Frhnmj2004/restaurant-admin/types"
)

type GroceryController struct {
	DB *gorm.DB
}

func NewGroceryController(db *gorm.DB) *GroceryController {
	return &GroceryController{DB: db}
}

func (ctrl *GroceryController) CreateGrocery(ctx *fiber.Ctx) error {
	grocery := models.Grocery{}

	err := ctx.BodyParser(&grocery)
	if err != nil {
		return helper.ErrorResponse(ctx, http.StatusBadRequest, "Invalid request body")
	}

	err = ctrl.DB.Create(&grocery).Error
	if err != nil {
		return helper.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to create grocery")
	}

	ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "Grocery created successfully",
	})

	return nil
}

func (ctrl *GroceryController) GetAllGroceries(ctx *fiber.Ctx) error {
	allGroceries := &[]models.Grocery{}

	err := ctrl.DB.Find(allGroceries).Error
	if err != nil {
		return helper.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to get all groceries")
	}

	ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "All groceries retrieved successfully",
		"data":    allGroceries,
	})

	return nil
}

func (ctrl *GroceryController) GetGroceryByName(ctx *fiber.Ctx) error {
	name := ctx.Params("name")
	if name == "" {
		return helper.ErrorResponse(ctx, http.StatusBadRequest, "Name is required")
	}

	groceryModel := &models.Grocery{}

	err := ctrl.DB.Where("name =?", name).First(groceryModel).Error
	if err != nil {
		return helper.ErrorResponse(ctx, http.StatusNotFound, "Grocery not found")
	}

	ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "Grocery retrieved successfully",
		"data":    groceryModel,
	})
	return nil
}

func (ctrl *GroceryController) UpdateGrocery(ctx *fiber.Ctx) error {
	name := ctx.Params("name")
	if name == "" {
		return helper.ErrorResponse(ctx, http.StatusBadRequest, "Name is required")
	}

	grocery := &models.Grocery{}

	err := ctrl.DB.Where("name =?", name).First(grocery).Error
	if err != nil {
		return helper.ErrorResponse(ctx, http.StatusNotFound, "Grocery not found")
	}

	request := &types.UpdateGroceryRequest{}

	err = ctx.BodyParser(request)
	if err != nil {
		return helper.ErrorResponse(ctx, http.StatusBadRequest, "Invalid request body")
	}

	grocery.Quantity += request.Quantity

	err = ctrl.DB.Save(grocery).Error
	if err != nil {
		return helper.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to update grocery")
	}

	ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "Grocery updated successfully",
		"data":    grocery,
	})

	return nil
}
