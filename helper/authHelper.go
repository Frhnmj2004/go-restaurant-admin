package helper

import (
	"log"

	"github.com/Frhnmj2004/restaurant-admin/models"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func CheckUserType(ctx *fiber.Ctx, role string) (err error) {
	return nil
}

func MatchUserTypetoUserID(ctx *fiber.Ctx, userID string) (err error) {
	return nil
}

func VerifyPassword(userpassword string, providedPassword string) (bool, string) {
	err := bcrypt.CompareHashAndPassword([]byte(providedPassword), []byte(userpassword))
	check := true
	msg := ""

	if err != nil {
		msg = "email or password is incorrect"
		check = false
	}
	return check, msg
}

func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Panic(err)
	}
	return string(bytes)
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
