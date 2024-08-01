package controllers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"github.com/Frhnmj2004/restaurant-admin/helper"
	"github.com/Frhnmj2004/restaurant-admin/models"
	"github.com/Frhnmj2004/restaurant-admin/types"
)

type OrderController struct {
	DB *gorm.DB
}

func NewOrderController(db *gorm.DB) *OrderController {
	return &OrderController{DB: db}
}

func (ctrl *OrderController) GetAllOrders(ctx *fiber.Ctx) error {
	allOrders := &[]models.Order{}

	err := ctrl.DB.Find(allOrders).Error
	if err != nil {
		return helper.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to get all orders")
	}

	ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "All orders retrieved successfully",
		"data":    allOrders,
	})

	return nil
}

func (ctrl *OrderController) GetOrderByID(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	if id == "" {
		return helper.ErrorResponse(ctx, http.StatusBadRequest, "Missing order ID")
	}

	order := &models.Order{}

	err := ctrl.DB.Where("id =?", id).First(order).Error
	if err != nil {
		return helper.ErrorResponse(ctx, http.StatusNotFound, "Order not found")
	}

	ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "Order retrieved successfully",
		"data":    order,
	})

	return nil
}

func (ctrl *OrderController) PlaceOrder(ctx *fiber.Ctx) error {
	request := &types.PlaceOrderRequest{}

	err := ctx.BodyParser(request)
	if err != nil {
		return helper.ErrorResponse(ctx, http.StatusBadRequest, "Invalid request body")
	}

	// debug
	//fmt.Printf("Placing order: %+v\n", request)

	foodItem := &models.FoodItem{}

	err = ctrl.DB.Preload("Ingredients.Grocery").Where("name =?", request.FoodItemName).First(foodItem).Error
	if err != nil {
		return helper.ErrorResponse(ctx, http.StatusNotFound, "Food item not found")
	}

	// debug
	//fmt.Printf("Food item found: %+v\n", foodItem)

	tx := ctrl.DB.Begin()

	revenue := &models.Revenue{}
	tx.First(revenue)

	revenue.TotalRevenue += foodItem.Price * float64(request.Quantity)

	err = tx.Save(revenue).Error
	if err != nil {
		tx.Rollback()
		return helper.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to update revenue")
	}

	// debug
	//fmt.Printf("lenght: %d\n", len(foodItem.Ingredients))

	for _, ingredient := range foodItem.Ingredients {
		grocery := &models.Grocery{}

		// debug
		//fmt.Printf("Grocery: %+v\n", ingredient)

		err := tx.Where("name =?", ingredient.Grocery.Name).First(grocery).Error
		if err != nil {
			tx.Rollback()
			return helper.ErrorResponse(ctx, http.StatusNotFound, "Grocery not found :"+ingredient.Grocery.Name)
		}

		requiredQuantity := ingredient.Quantity * float64(request.Quantity)

		if grocery.Quantity < requiredQuantity {
			tx.Rollback()
			return helper.ErrorResponse(ctx, http.StatusUnprocessableEntity, "Not enough quantity of "+ingredient.Grocery.Name+" for this order")
		}

		grocery.Quantity -= requiredQuantity

		err = tx.Save(grocery).Error
		if err != nil {
			tx.Rollback()
			return helper.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to update grocery")
		}
	}

	order := &models.Order{
		FoodItemID: foodItem.ID,
		Quantity:   request.Quantity,
		TotalPrice: foodItem.Price * float64(request.Quantity),
	}

	err = tx.Create(order).Error
	if err != nil {
		tx.Rollback()
		return helper.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to place order")
	}

	err = tx.Commit().Error
	if err != nil {
		return helper.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to commit transaction")
	}

	ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "Order placed successfully",
		"data":    order,
	})

	return nil
}
