package routes

import (
	"github.com/Frhnmj2004/restaurant-admin/controllers"
	"github.com/Frhnmj2004/restaurant-admin/middleware"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func FoodItemRoutes(app *fiber.App, db *gorm.DB) {
	foodItemCtrl := controllers.NewFoodItemController(db)
	//authenticate := middleware.Authenticate()

	foodItemGroup := app.Group("/api/fooditems", middleware.Authenticate)
	foodItemGroup.Post("/", foodItemCtrl.CreateFoodItem)
	foodItemGroup.Get("/", foodItemCtrl.GetAllFoodItems)
	foodItemGroup.Get("/:name", foodItemCtrl.GetFoodItemByName)
	foodItemGroup.Delete("/", foodItemCtrl.DeleteAllFoodItems)

}
