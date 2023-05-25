package routes

import (
	"praktikum/controllers"
	m "praktikum/middlewares"
	"praktikum/repository/database"
	"praktikum/usecase"
	"praktikum/util"

	"github.com/go-playground/validator"
	"github.com/labstack/echo"
	mid "github.com/labstack/echo/middleware"
	"gorm.io/gorm"
)

func NewRoute(e *echo.Echo, db *gorm.DB) {
	// Middleware
	m.LogMiddleware(e)
	e.Pre(mid.RemoveTrailingSlash())

	// Interface
	userRepository := database.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository)

	authRepository := database.NewAuthRepository(db)
	authUsecase := usecase.NewAuthUsecase(authRepository, userRepository)

	authController := controllers.NewAuthController(userUsecase, authUsecase)

	itemRepository := database.NewItemRepository(db)
	itemUsecase := usecase.NewItemUsecase(itemRepository)
	itemController := controllers.NewItemController(itemUsecase)

	categoryRepository := database.NewCategoryRepository(db)
	categoryUsecase := usecase.NewCategoryUsecase(categoryRepository)
	categoryController := controllers.NewCatergoryController(categoryUsecase)

	//validator
	e.Validator = &util.CustomValidator{Validator: validator.New()}

	// Routesres
	e.POST("/register", authController.RegisterController)
	e.POST("/login", authController.LoginController)

	// Routes Item
	i := e.Group("/items", m.IsLoggedIn)
	i.GET("", itemController.GetAllItemController)
	i.GET("", itemController.GetItemByNameController)
	i.GET("/:id", itemController.GetItemByIdController)
	i.GET("/category/:category_id", itemController.GetItemByCategoryIdController)
	i.POST("", itemController.CreateItemController)
	i.PUT("/:id", itemController.UpdateItemByIdController)
	i.DELETE("/:id", itemController.DeleteItemByIdController)

	// Routes Category
	c := e.Group("category", m.IsLoggedIn)
	c.GET("", categoryController.GetAllCatergoryController)
	c.GET("/:id", categoryController.GetCategoryByIdController)
	c.POST("", categoryController.CreateCategoryController)

}
