package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type User struct {
	Id       int    `json:"id" form:"id"`
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

var users []User
var Base_URL = "http://localhost:7500"

// -------------------- controller --------------------

// get all users
func GetUsersController(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success get all users",
		"users":    users,
	})
}

// get user by id
func GetUserController(c echo.Context) error {
	rex := map[string]interface{}{
		"messages": "users not found",
		"users":    nil,
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		rex["messages"] = err
		return c.JSON(http.StatusBadRequest, rex)
	}

	for _, val := range users {
		if val.Id == id {
			rex["messages"] = "user found"
			rex["users"] = val
		}
	}
	return c.JSON(http.StatusOK, rex)
}

// delete user by id
func DeleteUserController(c echo.Context) error {
	rex := map[string]interface{}{
		"messages": "users not found",
		"users":    nil,
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		rex["messages"] = err
		return c.JSON(http.StatusBadRequest, rex)
	}

	for idx := range users {
		if users[idx].Id == id {
			rex["users"] = users[idx]

			copy(users[idx:], users[idx+1:]) // Shift a[i+1:] left one index.
			users[len(users)-1] = User{}     // Erase last element (write zero value).
			users = users[:len(users)-1]

			rex["messages"] = "user deleted"

		}
	}

	return c.JSON(http.StatusOK, rex)
}

// update user by id
func UpdateUserController(c echo.Context) error {
	rex := map[string]interface{}{
		"messages": "users not found",
		"users":    nil,
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		rex["messages"] = err
		return c.JSON(http.StatusBadRequest, rex)
	}

	for idx := range users {
		if users[idx].Id == id {
			users[idx].Name = c.FormValue("name")
			users[idx].Email = c.FormValue("email")
			users[idx].Password = c.FormValue("password")

			rex["messages"] = "success update users"
			rex["users"] = users[idx]

			break
		}
	}

	return c.JSON(http.StatusOK, rex)
}

// create new user
func CreateUserController(c echo.Context) error {
	// binding data
	user := User{}
	c.Bind(&user)

	if len(users) == 0 {
		user.Id = 1
	} else {
		newId := users[len(users)-1].Id + 1
		user.Id = newId
	}
	users = append(users, user)
	return c.JSON(http.StatusCreated, map[string]interface{}{
		"messages": "success create user",
		"user":     user,
	})
}

// ---------------------------------------------------
func main() {
	e := echo.New()
	// routing with query parameter
	e.GET("/users", GetUsersController)
	e.GET("/users/:id", GetUserController)
	e.POST("/users", CreateUserController)
	e.PUT("/users/:id", UpdateUserController)
	e.DELETE("/users/:id", DeleteUserController)

	// start the server, and log if it fails
	e.Logger.Fatal(e.Start(":7500"))
}
