package routes

import (
	"middleware/constants"
	"middleware/controller"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {
	// create a new echo instance
	e := echo.New()
	jwtmiddleware := middleware.JWT([]byte(constants.SECRET_JWT))
	// Route / to handler function
	e.GET("/users", controller.GetUsersController, jwtmiddleware)
	e.GET("/users/:id", controller.GetUserController, jwtmiddleware)
	e.POST("/users", controller.CreateUserController)
	e.POST("/users/login", controller.LoginUserController)
	e.DELETE("/users/:id", controller.DeleteUserController, jwtmiddleware)
	e.PUT("/users/:id", controller.UpdateUserController)

	e.GET("/books", controller.GetBooksController, jwtmiddleware)
	e.GET("/books/:id", controller.GetBookController, jwtmiddleware)
	e.POST("/books", controller.CreateBookController, jwtmiddleware)
	e.DELETE("/books/:id", controller.DeleteBookController, jwtmiddleware)
	e.PUT("/books/:id", controller.UpdateBookController)

	return e
}
