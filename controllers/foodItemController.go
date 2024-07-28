package controllers

import (
	"net/http"

	"github.com/Frhnmj2004/restaurant-admin/helper"
	"github.com/Frhnmj2004/restaurant-admin/models"
	"github.com/Frhnmj2004/restaurant-admin/types"
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
	request := &types.CreateFoodItemRequest{}

	err := ctx.BodyParser(request)
	if err != nil {
		return helper.ErrorResponse(ctx, http.StatusBadRequest, "Invalid request body")
	}

	ingredients := make([]models.Ingredient, len(request.Ingredients))

	grocery := &models.Grocery{}

	for _, ing := range request.Ingredients {
		err := ctrl.DB.Where("name =?", ing.GroceryName).First(grocery).Error
		if err != nil {
			return helper.ErrorResponse(ctx, http.StatusNotFound, "Grocery not found :"+ing.GroceryName)
		}

		ingredients = append(ingredients, models.Ingredient{
			GroceryID:   grocery.ID,
			GroceryName: grocery.Name,
			Quantity:    ing.Quantity,
		})
	}

	foodItem := &models.FoodItem{
		Name:        request.Name,
		Price:       request.Price,
		Ingredients: ingredients,
	}

	err = ctrl.DB.Create(foodItem).Error
	if err != nil {
		return helper.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to create food item")
	}

	ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "Food item created successfully",
		"data":    foodItem,
	})

	return nil
}

func (ctrl *FoodItemController) GetFoodItemByName(ctx *fiber.Ctx) error {
	name := ctx.Params("name")
	if name == "" {
		return helper.ErrorResponse(ctx, http.StatusBadRequest, "Missing food item name")
	}

	foodItemModel := &models.FoodItem{}

	err := ctrl.DB.Where("name =?", name).First(foodItemModel).Error
	if err != nil {
		return helper.ErrorResponse(ctx, http.StatusNotFound, "Food item not found")
	}

	ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "Food item retrieved successfully",
		"data":    foodItemModel,
	})
	return nil
}

func (ctrl *FoodItemController) GetAllFoodItems(ctx *fiber.Ctx) error {
	allFoodItems := &[]models.FoodItem{}

	err := ctrl.DB.Find(allFoodItems).Error
	if err != nil {
		return helper.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to get all food items")
	}

	ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "All food items retrieved successfully",
		"data":    allFoodItems,
	})
	return nil
}
