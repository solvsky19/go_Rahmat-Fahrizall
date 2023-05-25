package controllers

import (
	"praktikum/models/payload"
	"praktikum/usecase"

	"github.com/labstack/echo"
)

type AuthController interface{}

type authController struct {
	userUsecase usecase.UserUsecase
	authUsecase usecase.AuthUsecase
}

func NewAuthController(
	userUsecase usecase.UserUsecase,
	authUsecase usecase.AuthUsecase,
) *authController {
	return &authController{
		userUsecase,
		authUsecase,
	}
}

func (aut *authController) LoginController(c echo.Context) error {
	reg := payload.LoginRequest{}

	c.Bind(&reg)

	if err := c.Validate(reg); err != nil {
		return echo.NewHTTPError(400, "Field cannot be empty")
	}

	res, err := aut.authUsecase.LoginUser(&reg)
	if err != nil {
		return echo.NewHTTPError(400, err.Error())
	}

	return c.JSON(200, payload.Response{
		Message: "Success login",
		Data:    res,
	})
}

func (aut *authController) RegisterController(c echo.Context) error {
	reg := payload.RegisterRequest{}

	c.Bind(&reg)

	if err := c.Validate(reg); err != nil {
		return echo.NewHTTPError(400, err.Error())
	}

	err := aut.userUsecase.CreateUser(&reg)
	if err != nil {
		return echo.NewHTTPError(400, err.Error())
	}

	return c.JSON(200, payload.Response{
		Message: "Success create user",
		Data:    reg,
	})

}
