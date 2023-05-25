package controllers

import (
	"praktikum/models"
	"praktikum/models/payload"
	"praktikum/usecase"
	"strconv"

	"github.com/labstack/echo"
)

type CategoryController interface {
	GetAllCatergoryController(c echo.Context) error
	GetCategoryByIdController(c echo.Context) error
}

type categoryController struct {
	categoryUsecase usecase.CatergoryUsecase
}

func NewCatergoryController(categoryUsecase usecase.CatergoryUsecase) *categoryController {
	return &categoryController{categoryUsecase}
}

func (cc *categoryController) GetAllCatergoryController(c echo.Context) error {
	category, err := cc.categoryUsecase.GetAllCategory()
	if err != nil {
		return err
	}

	return c.JSON(200, payload.Response{
		Message: "Success get all item",
		Data:    category,
	})
}

func (cc *categoryController) GetCategoryByIdController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	category, err := cc.categoryUsecase.GetCategoryById(id)
	if err != nil {
		return echo.NewHTTPError(400, err.Error())
	}

	return c.JSON(200, payload.Response{
		Message: "Success get category by id",
		Data:    category,
	})
}

func (cc *categoryController) CreateCategoryController(c echo.Context) error {
	var cat models.CategoryItems

	c.Bind(&cat)

	err := cc.categoryUsecase.CreateCategory(&cat)
	if err != nil {
		return echo.NewHTTPError(400, "Failed to create category")
	}

	return c.JSON(200, payload.Response{
		Message: "Success create category",
		Data:    cat,
	})
}
