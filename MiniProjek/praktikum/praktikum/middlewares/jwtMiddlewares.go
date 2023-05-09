package middlewares

import (
	"net/http"
	"praktikum/constants"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

func CreateToken(userId int) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["userId"] = userId
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(constants.SECRET_JWT))

}

func ExtractTokenUserId(c echo.Context) int {
	authHeader := c.Request().Header.Get("Authorization")
	if authHeader == "" {
		return 0
	}

	tokenString := strings.TrimPrefix(authHeader, "Bearer ")

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(constants.SECRET_JWT), nil
	})
	if err != nil || !token.Valid {
		return 0
	}

	claims := token.Claims.(jwt.MapClaims)
	userId := claims["userId"].(int)
	return int(userId)
}

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			return echo.NewHTTPError(http.StatusUnauthorized, "missing Authorization header")
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(constants.SECRET_JWT), nil
		})
		if err != nil || !token.Valid {
			return echo.NewHTTPError(http.StatusUnauthorized, "invalid or expired token")
		}

		claims := token.Claims.(jwt.MapClaims)
		userId := claims["userId"].(float64)
		c.Set("userId", int(userId))

		return next(c)
	}
}
