package routes

import (
	"github.com/Frhnmj2004/restaurant-admin/controllers"
	"github.com/Frhnmj2004/restaurant-admin/middleware"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RevenueRoutes(app *fiber.App, db *gorm.DB) {
	revenueCtrl := controllers.NewRevenueController(db)
	//authenticate := middleware.Authenticate()

	revenueGroup := app.Group("/api/revenue", middleware.Authenticate)
	revenueGroup.Get("/", revenueCtrl.GetRevenue)
}
