package controllers

import (
	"net/http"

	"github.com/Frhnmj2004/restaurant-admin/helper"
	"github.com/Frhnmj2004/restaurant-admin/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
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
	user := &models.User{}

	// parsing request body 
	err := ctx.BodyParser(user)
	if err!= nil {
		return helper.ErrorResponse(ctx, http.StatusInternalServerError, "cannot parse request body")
	}

	// validate user input
	validate := validator.New()
	validationErr := validate.Struct(user)
	if validationErr!= nil {
        return helper.ErrorResponse(ctx, http.StatusBadRequest, "invalid request body")
    }

	// email phone uniqueness check
	exists, err := helper.UserExists(ctrl.DB, *user.Email, *user.Phone) 
	if err!= nil {
        return helper.ErrorResponse(ctx, http.StatusInternalServerError, "failed to check email")
    }
	if exists {
        return helper.ErrorResponse(ctx, http.StatusConflict, "email or phone already exists")
    }

	


	return nil
}
