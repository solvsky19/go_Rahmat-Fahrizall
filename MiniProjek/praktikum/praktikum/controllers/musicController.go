package controllers

import (
	"fmt"
	"net/http"
	"praktikum/lib/database"
	"praktikum/models"
	"strconv"

	"github.com/labstack/echo"
)

func GetAllMusicsController(c echo.Context) error {
	musics, err := database.GetAllMusics()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all Musics",
		"musics":  musics,
	})
}

func GetMusicByIDController(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid music ID")
	}

	music, err := database.GetMusicByID(uint(id))
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, fmt.Sprintf("Music with ID: %d Not Found", id))
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": fmt.Sprintf("success Get Music with ID: %d", id),
		"music":   music,
	})
}

func CreateMusicController(c echo.Context) error {
	var music models.Musics

	if err := c.Bind(&music); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	newMusic, err := database.CreateMusic(&music)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "success create new Music",
		"music":   newMusic,
	})
}

func UpdateMusicController(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid Music ID")
	}

	var music models.Musics
	if err := c.Bind(&music); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	updatedMusic, err := database.UpdateMusicByID(uint(id), &music)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": fmt.Sprintf("success Update Music with ID: %d", id),
		"music":   updatedMusic,
	})
}

func DeleteMusicController(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid Music ID")
	}

	err = database.DeleteMusicByID(uint(id))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": fmt.Sprintf("success Delete Music with ID: %d", id),
	})
}
