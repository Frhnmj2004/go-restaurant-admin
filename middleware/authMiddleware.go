package middleware

import (
	"net/http"

	"github.com/Frhnmj2004/restaurant-admin/helper"
	"github.com/gofiber/fiber/v2"
)

func Authenticate(ctx *fiber.Ctx) error {
	clientToken := ctx.Get("token")
	if clientToken == "" {
		return helper.ErrorResponse(ctx, http.StatusBadRequest, "No Authorization token provided")
	}

	claims, err := helper.ValidateToken(clientToken)
	if err != "" {
		return helper.ErrorResponse(ctx, http.StatusUnauthorized, "Invalid token")
	}

	ctx.Locals("user_id", claims.UserID)
	ctx.Locals("first_name", claims.FirstName)
	ctx.Locals("last_name", claims.LastName)
	ctx.Locals("email", claims.Email)
	ctx.Locals("user_type", claims.UserType)

	return ctx.Next()
}
