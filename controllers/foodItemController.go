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

	tx := ctrl.DB.Begin()

	foodItem := &models.FoodItem{
		Name:  request.Name,
		Price: request.Price,
	}

	err = tx.Create(foodItem).Error
	if err != nil {
		tx.Rollback()
		return helper.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to create food item")
	}

	for _, ing := range request.Ingredients {

		// debug
		//fmt.Printf("Creating ingredient: %+v\n", ing)

		grocery := &models.Grocery{}

		err := tx.Where("name = ?", ing.GroceryName).First(grocery).Error
		if err != nil {
			tx.Rollback()
			return helper.ErrorResponse(ctx, http.StatusNotFound, "Grocery not found: "+ing.GroceryName)
		}

		// debug
		//fmt.Printf("Grocery found: %+v\n", grocery)

		ingredient := models.Ingredient{
			FoodItemID: foodItem.ID,
			GroceryID:  grocery.ID,
			Quantity:   ing.Quantity,
			//Grocery:    *grocery,
			//FoodItem:   *foodItem,
		}

		// debug
		//fmt.Printf("Ingredient created: %+v\n\n", ingredient)

		err = tx.Create(&ingredient).Error
		if err != nil {
			tx.Rollback()
			return helper.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to create ingredient")
		}

		foodItem.Ingredients = append(foodItem.Ingredients, ingredient)

		err = tx.Save(foodItem).Error
		if err != nil {
			tx.Rollback()
			return helper.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to update food item")
		}

		// debug
		//fmt.Printf("Food item updated: %+v\n\n", foodItem)
	}

	tx.Commit()

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

func (ctrl *FoodItemController) DeleteAllFoodItems(ctx *fiber.Ctx) error {
	tx := ctrl.DB.Begin()

	err := tx.Exec("DELETE FROM orders WHERE food_item_id IS NOT NULL").Error
	if err != nil {
		tx.Rollback()
		return helper.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to delete all orders related to food items")
	}

	err = tx.Exec("DELETE FROM fooditem_ingredients").Error
	if err != nil {
		tx.Rollback()
		return helper.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to delete all food item ingredients")
	}

	err = tx.Exec("DELETE FROM ingredients").Error
	if err != nil {
		tx.Rollback()
		return helper.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to delete all ingredients")
	}

	err = tx.Exec("DELETE FROM food_items").Error
	if err != nil {
		tx.Rollback()
		return helper.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to delete all food items")
	}

	err = tx.Exec("ALTER SEQUENCE food_items_id_seq RESTART WITH 1").Error
	if err != nil {
		tx.Rollback()
		return helper.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to reset food item ID sequence")
	}

	err = tx.Commit().Error
	if err != nil {
		return helper.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to delete all food items and associated data")
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "All food items, related orders, and associated data deleted successfully, and ID sequence reset",
	})
}
