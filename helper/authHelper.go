package helper

import "github.com/gofiber/fiber/v2"

func CheckUserType(ctx *fiber.Ctx, role string) (err error) {
	return nil
}

func MatchUserTypetoUserID(ctx *fiber.Ctx, userID string) (err error) {
	return nil
}

func VerifyPassword(userpassword string, providedPassword string) (bool, string) {
	return true, ""
}

func HashPassword(password string) string {
	return ""
}
