package controllers

import (
	"fmt"
	"praktikum/models/payload"
	"praktikum/usecase"
	"strconv"

	"github.com/labstack/echo"
)

type ItemController interface {
	GetAllItemController(c echo.Context) error
	GetItemByIdController(c echo.Context) error
	GetItemByCategoryIdController(c echo.Context) error
	CreateItemController(c echo.Context) error
	UpdateItemByIdController(c echo.Context)
	DeleteItemByIdController(c echo.Context) error
}

type itemController struct {
	itemUsecase usecase.ItemUsecase
}

func NewItemController(itemUsecase usecase.ItemUsecase) *itemController {
	return &itemController{itemUsecase}
}

func (ic *itemController) GetAllItemController(c echo.Context) error {
	item, err := ic.itemUsecase.GetAllitems()
	if err != nil {
		return err
	}

	return c.JSON(200, payload.Response{
		Message: "Success get all item",
		Data:    item,
	})
}

func (ic *itemController) GetItemByIdController(c echo.Context) error {
	id := c.Param("id")

	item, err := ic.itemUsecase.GetItemById(id)
	if err != nil {
		return echo.NewHTTPError(400, err.Error())
	}

	return c.JSON(200, payload.Response{
		Message: fmt.Sprintf("item %s", item.Name),
		Data:    item,
	})
}

func (ic *itemController) GetItemByNameController(c echo.Context) error {
	name := c.QueryParam("name")

	item, err := ic.itemUsecase.GetItemByName(name)
	if err != nil {
		return err
	}

	return c.JSON(200, payload.Response{
		Message: "Success get items",
		Data:    item,
	})
}

func (ic *itemController) GetItemByCategoryIdController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("category_id"))

	item, err := ic.itemUsecase.GetItemByCategoryId(id)
	if err != nil {
		return echo.NewHTTPError(400, err.Error())
	}

	return c.JSON(200, payload.Response{
		Message: "Success get item by id category",
		Data:    item,
	})
}

func (ic *itemController) CreateItemController(c echo.Context) error {
	req := payload.CreateItem{}

	c.Bind(&req)

	if err := c.Validate(req); err != nil {
		return echo.NewHTTPError(400, "Field Cannot be empty")
	}

	item, err := ic.itemUsecase.CreateItem(&req)
	if err != nil {
		return echo.NewHTTPError(400, err.Error())
	}

	return c.JSON(200, payload.Response{
		Message: "Success Create Item",
		Data:    item,
	})
}

func (ic *itemController) UpdateItemByIdController(c echo.Context) error {
	req := payload.UpdateItem{}

	id := c.Param("id")

	c.Bind(&req)

	if err := c.Validate(req); err != nil {
		return echo.NewHTTPError(400, "Field cannot be empty")
	}

	item, err := ic.itemUsecase.UpdateItemById(id, &req)
	if err != nil {
		return echo.NewHTTPError(400, err.Error())
	}

	return c.JSON(200, payload.Response{
		Message: "Success Update item",
		Data:    item,
	})
}

func (ic *itemController) DeleteItemByIdController(c echo.Context) error {
	id := c.Param("id")

	_, err := ic.itemUsecase.DeleteItemById(id)
	if err != nil {
		return c.JSON(400, err.Error())
	}

	return c.JSON(200, "Success delete item")
}
