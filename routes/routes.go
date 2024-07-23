package routes

import (
	"github.com/Frhnmj2004/restaurant-admin/controllers"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{DB: db}
}
func (r *Repository) SetupRoutes(app *fiber.App) {
	groceryCtrl := controllers.NewGroceryController(r.DB)
	foodItemCtrl := controllers.NewFoodItemController(r.DB)
	orderCtrl := controllers.NewOrderController(r.DB)

	groceryGroup := app.Group("/api/groceries")
	groceryGroup.Post("/", groceryCtrl.CreateGrocery)
	groceryGroup.Get("/", groceryCtrl.GetAllGroceries)
	groceryGroup.Get("/:id", groceryCtrl.GetGroceryByID)
	groceryGroup.Put("/:id", groceryCtrl.UpdateGrocery)

	foodItemGroup := app.Group("/api/fooditems")
	foodItemGroup.Post("/", foodItemCtrl.CreateFoodItem)
	foodItemGroup.Get("/", foodItemCtrl.GetAllFoodItems)
	foodItemGroup.Get("/:id", foodItemCtrl.GetFoodItemByID)

	orderGroup := app.Group("/api/orders")
	orderGroup.Post("/", orderCtrl.PlaceOrder)
	orderGroup.Get("/", orderCtrl.GetAllOrders)
	orderGroup.Get("/:id", orderCtrl.GetOrderByID)
}
