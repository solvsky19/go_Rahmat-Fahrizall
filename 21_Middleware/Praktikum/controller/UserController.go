package controller

import (
	"middleware/config"
	"middleware/lib/database"
	"middleware/model"
	"net/http"

	"strconv"

	"github.com/labstack/echo/v4"
)

func LoginUserController(c echo.Context) error {
	user := model.User{}
	c.Bind(&user)

	users, e := database.LoginUsers(&user)
	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success login",
		"users":  users,
	})
}

func GetUserDetailControllers(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	users, err := database.GetDetailUsers((id))

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"users":  users,
	})
}

// get all user
func GetUsersController(c echo.Context) error {
	var users []model.User

	if err := config.DB.Find(&users).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all users",
		"users":   users,
	})
}

// get user by id
func GetUserController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": "parameter salah",
		})
	}
	var user model.User
	if err := config.DB.First(&user, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"status": "user not found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":    "berhasil",
		"data_user": user,
	})
}

func CreateUserController(c echo.Context) error {
	user := model.User{}
	c.Bind(&user)

	if err := config.DB.Save(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new user",
		"user":    user,
	})
}

// delete user by id
func DeleteUserController(c echo.Context) error {
	// your solution here
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": "parameter salah",
		})
	}

	var user model.User
	if err := config.DB.Where("id = ?", id).First(&user).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"status": "user not available",
		})
	}

	if err := config.DB.Delete(&user).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"status": "gagal hapus data"})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":    "berhasil menghapus data",
		"data_user": user,
	})

}

func UpdateUserController(c echo.Context) error {
	bady := new(model.User)

	if err := c.Bind(bady); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": "parameter salah",
		})
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": "invalid id parameter",
		})
	}

	var user model.User
	if err := config.DB.Where("id = ?", id).First(&user).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"status": "user not available",
		})
	}

	user.Name = bady.Name
	user.Email = bady.Email
	user.Password = bady.Password

	if err := config.DB.Save(&user).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"status": "gagal update data"})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "satus berhasil update data",
		"staus":  user,
	})
}
