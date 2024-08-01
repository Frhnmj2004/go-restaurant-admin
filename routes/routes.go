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
	revenueCtrl := controllers.NewRevenueController(r.DB)

	groceryGroup := app.Group("/api/groceries")
	groceryGroup.Post("/", groceryCtrl.CreateGrocery)
	groceryGroup.Get("/", groceryCtrl.GetAllGroceries)
	groceryGroup.Get("/:name", groceryCtrl.GetGroceryByName)
	groceryGroup.Put("/:name", groceryCtrl.UpdateGrocery)

	foodItemGroup := app.Group("/api/fooditems")
	foodItemGroup.Post("/", foodItemCtrl.CreateFoodItem)
	foodItemGroup.Get("/", foodItemCtrl.GetAllFoodItems)
	foodItemGroup.Get("/:name", foodItemCtrl.GetFoodItemByName)
	foodItemGroup.Delete("/", foodItemCtrl.DeleteAllFoodItems)

	orderGroup := app.Group("/api/orders")
	orderGroup.Post("/", orderCtrl.PlaceOrder)
	orderGroup.Get("/", orderCtrl.GetAllOrders)
	orderGroup.Get("/:id", orderCtrl.GetOrderByID)

	revenueGroup := app.Group("/api/revenue")
	revenueGroup.Get("/", revenueCtrl.GetRevenue)

}
