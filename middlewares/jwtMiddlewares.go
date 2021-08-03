package middlewares

import (
	"project/constants"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

//create token with adding limit time
func CreateToken(userId int) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["userId"] = userId
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(constants.SECRET_JWT))
}

func ExtractTokenCustomerId(e echo.Context) int {
	customer := e.Get("user").(*jwt.Token)

	if customer.Valid {
		claims := customer.Claims.(jwt.MapClaims)
		customerId := int(claims["userId"].(float64))
		return customerId
	}
	return 0
}
