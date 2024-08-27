package helper

import (
	"log"
	"os"
	"time"

	"github.com/Frhnmj2004/restaurant-admin/types"
	"github.com/dgrijalva/jwt-go"
)

var SECRET_KEY string = os.Getenv("SECRET_KEY")

func GenerateAllTokens(details *types.SignedDetails) (signedTokens string, signedRefreshTokens string, err error) {
	claims := *details
	claims.StandardClaims = jwt.StandardClaims{
		ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(24)).Unix(),
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(SECRET_KEY))
	if err != nil {
		log.Panic(err)
		return "", "", err
	}

	refreshClaims := *details
	refreshClaims.StandardClaims = jwt.StandardClaims{
		ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(360)).Unix(),
	}
	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString([]byte(SECRET_KEY))
	if err != nil {
		log.Panic(err)
		return "", "", err
	}

	return token, refreshToken, nil
}
func ValidateToken(signedToken string) (claims *types.SignedDetails, msg string) {
	return nil, ""
}

func UpdateAllTokens(signedToken string) (string, error) {
	return "", nil
}
