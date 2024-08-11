package routes

import (
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
	GroceryRoutes(app, r.DB)
	FoodItemRoutes(app, r.DB)
	AuthRoutes(app, r.DB)
	RevenueRoutes(app, r.DB)
	OrderRoutes(app, r.DB)
}
