package helper

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"github.com/Frhnmj2004/restaurant-admin/models"
) 

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

func UserExists(DB *gorm.DB, email string, phone string) (bool, error) {
	user := &models.User{}
	err := DB.Where("email =? OR phone =?", email, phone).First(user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
