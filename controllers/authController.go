package controllers

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type AuthController struct {
	DB *gorm.DB
}

func NewAuthController(db *gorm.DB) *AuthController {
	return &AuthController{DB: db}
}

func (ctrl *AuthController) Login(ctx *fiber.Ctx) error {
	return nil
}

func (ctrl *AuthController) Signup(ctx *fiber.Ctx) error {
	return nil
}
