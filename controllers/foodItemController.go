package controllers

import (
	"net/http"

	"github.com/Frhnmj2004/restaurant-admin/helper"
	"github.com/Frhnmj2004/restaurant-admin/models"
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
	type IngredientRequest struct {
		GroceryName string  `json:"groceryname"`
		Quantity    float64 `json:"quantity"`
	}

	type CreateFoodItemRequest struct {
		Name        string              `json:"name"`
		Price       float64             `json:"price"`
		Ingredients []IngredientRequest `json:"ingredients"`
	}

	request := &CreateFoodItemRequest{}
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

	return nil
}

func (ctrl *FoodItemController) GetAllFoodItems(ctx *fiber.Ctx) error {

	return nil
}
