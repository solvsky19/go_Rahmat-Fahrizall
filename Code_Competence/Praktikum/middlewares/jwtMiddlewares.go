package middleware

import (
	"praktikum/constants"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	mid "github.com/labstack/echo/middleware"
)

var IsLoggedIn = mid.JWTWithConfig(mid.JWTConfig{
	SigningMethod: "HS256",
	SigningKey:    []byte(constants.SECRET_JWT),
})

// Create Token Jwt
func CreateToken(userId int) (string, error) {
	claims := jwt.MapClaims{}
	claims["user_id"] = userId
	claims["authorized"] = true
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(constants.SECRET_JWT))
}