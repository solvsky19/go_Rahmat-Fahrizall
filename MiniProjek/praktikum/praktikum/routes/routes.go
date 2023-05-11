package routes

import (
	"praktikum/constants"
	"praktikum/controllers"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func New() *echo.Echo {
	// create a new echo instance
	e := echo.New()
	// auth.Use(echojwt.JWT([]byte(constants.SECRET_JWT)))

	//Auth
	e.POST("/login", controllers.LoginUsersController)
	e.POST("/register/user", controllers.RegisterUserController)
	// e.POST("/login", controllers.LoginAdminController)
	// e.POST("/register/admin", controllers.RegisterAdminController)
	//users

	auth := e.Group("", middleware.JWT([]byte(constants.SECRET_JWT)))
	// auth.Use(middlewares.AuthMiddleware)
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
	auth.PUT("/tikets/:id", controllers.UpdateTiketController)
	auth.DELETE("/tikets/:id", controllers.DeleteTiketController)

	//Tikets
	auth.GET("/InformasiAcaras", controllers.GetAllTiketsController)
	auth.GET("/InformasiAcaras/:id", controllers.GetTiketByIdController)
	auth.POST("/InformasiAcaras", controllers.CreateTiketController)
	auth.PUT("/InformasiAcaras/:id", controllers.UpdateTiketController)
	auth.DELETE("/InformasiAcaras/:id", controllers.DeleteTiketController)
	// start the server, and log if it fails
	return e
}
