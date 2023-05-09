package controllers

import (
	"fmt"
	"net/http"
	"praktikum/lib/database"
	"praktikum/models"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetAllInformasiAcarasController(c echo.Context) error {
	InformasiAcaras, err := database.GetAllInformasiAcaras()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":         "success get all InformasiAcaras",
		"InformasiAcaras": InformasiAcaras,
	})
}

// get InformasiAcara by id
func GetInformasiAcaraByIdController(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid InformasiAcara ID")
	}

	InformasiAcara, err := database.GetInformasiAcaraById(uint(id))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": fmt.Sprintf("Tiket with ID: %d Not Found", id),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":        fmt.Sprintf("success Get InformasiAcara with ID: %d", id),
		"InformasiAcara": InformasiAcara,
	})
}

// create new Tikets
func CreateInformasiAcaraController(c echo.Context) error {
	var InformasiAcara models.InformasiAcaras

	if err := c.Bind(&InformasiAcara); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	newInformasiAcara, err := database.CreateInformasiAcaras(&InformasiAcara)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message":        "success create new InformasiAcara",
		"InformasiAcara": newInformasiAcara,
	})
}

// update Tiket by id
func UpdateInformasiAcaraController(c echo.Context) error {
	var InformasiAcara models.InformasiAcaras

	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid InformasiAcara ID")
	}

	if err := c.Bind(&InformasiAcara); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	updateInformasiAcara, err := database.UpdateInformasiAcaraById(uint(id), &InformasiAcara)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":        fmt.Sprintf("success Update InformasiAcara with ID: %d", id),
		"InformasiAcara": updateInformasiAcara,
	})
}

// delete Tiket by id
func DeleteInformasiAcaraController(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid InformasiAcara ID")
	}

	err = database.DeleteInformasiAcaraById(uint(id))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": fmt.Sprintf("success Delete InformasiAcara with ID: %d", id),
	})
}
