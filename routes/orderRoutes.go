package routes

import (
	"github.com/Frhnmj2004/restaurant-admin/controllers"
	"github.com/Frhnmj2004/restaurant-admin/middleware"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func OrderRoutes(app *fiber.App, db *gorm.DB) {
	orderCtrl := controllers.NewOrderController(db)
	//authenticate := middleware.Authenticate()

	orderGroup := app.Group("/api/orders", middleware.Authenticate)
	orderGroup.Post("/", orderCtrl.PlaceOrder)
	orderGroup.Get("/", orderCtrl.GetAllOrders)
	orderGroup.Get("/:id", orderCtrl.GetOrderByID)

}
