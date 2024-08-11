package routes

import (
	"github.com/Frhnmj2004/restaurant-admin/controllers"
	"github.com/Frhnmj2004/restaurant-admin/middleware"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GroceryRoutes(app *fiber.App, db *gorm.DB) {
	groceryCtrl := controllers.NewGroceryController(db)
	//authenticate := middleware.Authenticate()

	groceryGroup := app.Group("/api/groceries", middleware.Authenticate)
	groceryGroup.Post("/", groceryCtrl.CreateGrocery)
	groceryGroup.Get("/", groceryCtrl.GetAllGroceries)
	groceryGroup.Get("/:name", groceryCtrl.GetGroceryByName)
	groceryGroup.Put("/:name", groceryCtrl.UpdateGrocery)

}
