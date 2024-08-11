package routes

import (
	"github.com/Frhnmj2004/restaurant-admin/controllers"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func AuthRoutes(app *fiber.App, db *gorm.DB) {
	authCtrl := controllers.NewAuthController(db)

	authGroup := app.Group("/api/auth")
	authGroup.Post("/login", authCtrl.Login)
	authGroup.Post("/register", authCtrl.Signup)

}
