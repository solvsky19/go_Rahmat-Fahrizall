package controllers

import (
	"net/http"
	"praktikum/constants"
	"praktikum/lib/database"
	"praktikum/models"

	"github.com/labstack/echo"
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
	user.Role = constants.USER
	u, err := database.CreateUser(&user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "success create new user",
		"user":    u,
	})
}

// func RegisterAdminController(c echo.Context) error {
// 	req := models.Users{}

// 	c.Bind(&req)

// 	if err := c.Validate(&req); err != nil {
// 		return echo.NewHTTPError(400, "Field cannot be empty or Password must be 6 character")
// 	}
// 	req.Role = constants.ADMIN
// 	user, err := database.CreateUser(&req)

// 	if err != nil {
// 		return echo.NewHTTPError(400, err.Error())
// 	}

// 	return c.JSON(200, map[string]interface{}{
// 		"message": "Success Register",
// 		"users":   user,
// 	})

// }

// func LoginAdminController(c echo.Context) error {
// 	req := payload.Logins{}

// 	c.Bind(&req)

// 	if err := c.Validate(&req); err != nil {
// 		return echo.NewHTTPError(400, "Field cannot be empty")
// 	}

// 	res, err := database.LoginUsers(&req)

// 	if err != nil {
// 		return echo.NewHTTPError(400, "Invalid Email or Password")
// 	}

// 	return c.JSON(200, payload.Response{
// 		Message: "Success Login",
// 		Data:    res,
// 	})
// }
