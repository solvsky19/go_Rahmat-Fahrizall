package controllers

import (
	"fmt"
	"net/http"
	"praktikum/lib/database"
	"praktikum/models"
	"strconv"

	"github.com/labstack/echo"
)

func GetAllTiketsController(c echo.Context) error {
	tickets, err := database.GetAllTikets()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all Tikets",
		"tickets": tickets,
	})
}

// get Tiket by id
func GetTiketByIdController(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid tiket ID")
	}

	ticket, err := database.GetTiketById(uint(id))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": fmt.Sprintf("Tiket with ID: %d Not Found", id),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": fmt.Sprintf("success Get Tiket with ID: %d", id),
		"tiket":   ticket,
	})
}

// create new Tikets
func CreateTiketController(c echo.Context) error {
	var tiket models.Tikets

	if err := c.Bind(&tiket); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	newTiket, err := database.CreateTikets(&tiket)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "success create new Tiket",
		"tiket":   newTiket,
	})
}

// update Tiket by id
func UpdateTiketController(c echo.Context) error {
	var tiket models.Tikets

	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid Tiket ID")
	}

	if err := c.Bind(&tiket); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	updateTiket, err := database.UpdateTiketById(uint(id), &tiket)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": fmt.Sprintf("success Update Tiket with ID: %d", id),
		"tiket":   updateTiket,
	})
}

// delete Tiket by id
func DeleteTiketController(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid Tiket ID")
	}

	err = database.DeleteTiketById(uint(id))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": fmt.Sprintf("success Delete Tiket with ID: %d", id),
	})
}
