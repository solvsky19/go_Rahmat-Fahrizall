package routes

import (
	"praktikum/controllers"
	"praktikum/middlewares"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	// create a new echo instance
	e := echo.New()
	// auth.Use(echojwt.JWT([]byte(constants.SECRET_JWT)))

	//Auth
	e.POST("/login", controllers.LoginUsersController)
	e.POST("/register", controllers.RegisterUserController)

	//users

	auth := e.Group("")
	auth.Use(middlewares.AuthMiddleware)
	auth.GET("/users", controllers.GetAllUserControllers)
	auth.GET("/users/:id", controllers.GetUserByIdController)
	auth.PUT("/users/:id", controllers.UpdateUserController)
	auth.DELETE("/users/:id", controllers.DeleteUserController)

	//Musics
	auth.GET("/musics", controllers.GetAllMusicsController)
	auth.GET("/musics/:id", controllers.GetMusicByIDController)
	auth.POST("/musics", controllers.CreateMusicController)
	auth.PUT("/musics/:id", controllers.UpdateMusicController)
	auth.DELETE("/musics/:id", controllers.DeleteMusicController)

	//Tikets
	auth.GET("/tikets", controllers.GetAllTiketsController)
	auth.GET("/tikets/:id", controllers.GetTiketByIdController)
	auth.POST("/tikets", controllers.CreateTiketController)
	auth.PUT("/tikets/:id", controllers.UpdateMusicController)
	auth.DELETE("/tikets/:id", controllers.DeleteMusicController)

	// start the server, and log if it fails
	return e
}
