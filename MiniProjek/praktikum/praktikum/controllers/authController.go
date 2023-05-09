package controllers

import (
	"net/http"
	"praktikum/lib/database"
	"praktikum/models"

	"github.com/labstack/echo/v4"
)

func LoginUsersController(c echo.Context) error {
	user := models.Users{}
	if err := c.Bind(&user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request body")
	}
	if user.Email == "" || user.Password == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "email and password are required")
	}

	token, err := database.LoginUsers(&user)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "invalid email or password")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":   "success login",
		"JWT token": token,
	})
}

// create new user
func RegisterUserController(c echo.Context) error {
	var user models.Users

	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	u, err := database.CreateUser(&user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "success create new user",
		"user":    u,
	})
}
