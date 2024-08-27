package controllers

import (
	"errors"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/Frhnmj2004/restaurant-admin/helper"
	"github.com/Frhnmj2004/restaurant-admin/models"
	"github.com/Frhnmj2004/restaurant-admin/types"
	"github.com/go-playground/validator/v10"
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
	user := &models.User{}
	foundUser := &models.User{}

	// parsing the request body
	err := ctx.BodyParser(user)
	if err != nil {
		return helper.ErrorResponse(ctx, http.StatusBadRequest, "cannot parse the request body")
	}

	// find user by email
	err = ctrl.DB.Where("email = ?", user.Email).First(foundUser).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return helper.ErrorResponse(ctx, http.StatusNotFound, "the email Id incorrect")
		}
		return helper.ErrorResponse(ctx, http.StatusInternalServerError, "failed to find user")
	}

	// password verification
	passwordIsValid, msg := helper.VerifyPassword(*user.Password, *foundUser.Password)
	if !passwordIsValid {
		return helper.ErrorResponse(ctx, http.StatusUnauthorized, msg)
	}

	// generate JWT token
	signedDetails := &types.SignedDetails{
		Email:     *foundUser.Email,
		FirstName: *foundUser.FirstName,
		LastName:  *foundUser.LastName,
		UserID:    strconv.FormatUint(uint64(foundUser.ID), 10),
		UserType:  *foundUser.UserType,
	}
	token, refreshToken, err := helper.GenerateAllTokens(signedDetails)
	if err != nil {
		return helper.ErrorResponse(ctx, http.StatusInternalServerError, "failed to generate tokens")
	}

	// updating tokens in database
	foundUser.Token = &token
	foundUser.RefreshToken = &refreshToken

	err = ctrl.DB.Save(foundUser).Error
	if err != nil {
		return helper.ErrorResponse(ctx, http.StatusInternalServerError, "failed to update tokens")
	}

	// response with JWT tokens
	ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "Login successful",
		"data": fiber.Map{
			"access_token":  token,
			"refresh_token": refreshToken,
		},
	})
	return nil
}

func (ctrl *AuthController) Signup(ctx *fiber.Ctx) error {
	user := &models.User{}

	// parsing request body
	err := ctx.BodyParser(user)
	if err != nil {
		return helper.ErrorResponse(ctx, http.StatusInternalServerError, "cannot parse request body")
	}

	// validate user input
	validate := validator.New()
	validationErr := validate.Struct(user)
	if validationErr != nil {
		return helper.ErrorResponse(ctx, http.StatusBadRequest, "invalid request body")
	}

	// email phone uniqueness check
	exists, err := helper.UserExists(ctrl.DB, *user.Email, *user.Phone)
	if err != nil {
		return helper.ErrorResponse(ctx, http.StatusInternalServerError, "failed to check email")
	}
	if exists {
		return helper.ErrorResponse(ctx, http.StatusConflict, "email or phone already exists")
	}

	// logging the created time
	user.CreatedAt, err = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	if err != nil {
		log.Panic(err)
		return helper.ErrorResponse(ctx, http.StatusInternalServerError, "failed to parse created at")
	}

	user.UpdatedAt, err = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	if err != nil {
		log.Panic(err)
		return helper.ErrorResponse(ctx, http.StatusInternalServerError, "failed to parse updated at")
	}

	// Hash the users password
	password := helper.HashPassword(*user.Password)
	user.Password = &password

	// Generate Tokens
	signedDetails := &types.SignedDetails{
		Email:     *user.Email,
		FirstName: *user.FirstName,
		LastName:  *user.LastName,
		UserID:    strconv.FormatUint(uint64(user.ID), 10),
		UserType:  *user.UserType,
	}
	token, refreshToken, err := helper.GenerateAllTokens(signedDetails)
	if err != nil {
		log.Panic(err)
		return helper.ErrorResponse(ctx, http.StatusInternalServerError, "failed to generate tokens")
	}

	// updating tokens in database
	user.Token = &token
	user.RefreshToken = &refreshToken

	// save user to database
	err = ctrl.DB.Create(user).Error
	if err != nil {
		return helper.ErrorResponse(ctx, http.StatusInternalServerError, "failed to create user")
	}

	ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "User Created Successfully",
		"user":    user,
	})
	return nil
}
