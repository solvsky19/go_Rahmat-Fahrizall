package controllers

import (
	"fmt"
	"net/http"
	"praktikum/lib/database"

	"praktikum/models"
	"praktikum/models/payload"
	"strconv"

	"github.com/labstack/echo"
)

func CreateOrderController(c echo.Context) error {

	var orders models.Order

	if err := c.Bind(&orders); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	fmt.Println(orders.ID, orders.MusicID, orders.UserID, orders.TiketID)

	newOrder, err := database.CreateOrder(orders)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "success create new Order",
		"order":   newOrder,
	})

}

func GetOrderByIdControllers(c echo.Context) error {
	// id := middlewares.ExtractTokenUserId(c)
	id := c.Param("id")

	idx, _ := strconv.Atoi(id)
	orders, err := database.GetOrderByID(idx)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, payload.Responses{
		Message: "success get order",
		Data:    orders,
	})
}

func UpdateOrderController(c echo.Context) error {
	var order models.Order

	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid Order ID")
	}

	if err := c.Bind(&order); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	updateOrders, err := database.UpdateOrderByID(uint(id), &order)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": fmt.Sprintf("success Update Order with ID: %d", id),
		"tiket":   updateOrders,
	})
}

// delete Tiket by id
func DeleteOrderController(c echo.Context) error {
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
