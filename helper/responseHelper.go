package helper

import (
	"github.com/gofiber/fiber/v2"
)

func ErrorResponse(ctx *fiber.Ctx, status int, message string) error {
	return ctx.Status(status).JSON(fiber.Map{"error": message})
}
