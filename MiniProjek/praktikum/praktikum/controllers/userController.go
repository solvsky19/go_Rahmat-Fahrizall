package controllers

import (
	"fmt"
	"net/http"
	"praktikum/lib/database"
	"praktikum/models"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetAllUserControllers(c echo.Context) error {
	users, err := database.GetAllUsers()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all users",
		"users":   users,
	})
}

// get user by id
func GetUserByIdController(c echo.Context) error {
	// your solution here
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid user ID")
	}

	user, err := database.GetUserById(uint(id))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages": fmt.Sprintf("users with ID: %d Not Found", id),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": fmt.Sprintf("success Get user with ID: %d", id),
		"user":     user,
	})
}

// update user by id
func UpdateUserController(c echo.Context) error {
	// your solution here
	var user models.Users

	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid user ID")
	}

	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	updateUser, err := database.UpdateUserById(uint(id), &user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": fmt.Sprintf("success Update user with ID: %d", id),
		"user":     updateUser,
	})
}

// delete user by id
func DeleteUserController(c echo.Context) error {
	// your solution here
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid user ID")
	}

	err = database.DeleteUserById(uint(id))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": fmt.Sprintf("success Delete user with ID: %d", id),
	})
}
