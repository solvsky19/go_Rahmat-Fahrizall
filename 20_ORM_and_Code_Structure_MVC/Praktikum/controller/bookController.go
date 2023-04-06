package controller

import (
	"net/http"
	"praktikum/config"
	"praktikum/model"
	"strconv"

	"github.com/labstack/echo/v4"
)

// get all user
func GetBooksController(c echo.Context) error {
	var books []model.Book

	if err := config.DB.Find(&books).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all books",
		"books":   books,
	})
}

// get user by id
func GetBookController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": "parameter salah",
		})
	}
	var book model.Book
	if err := config.DB.First(&book, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"status": "book not found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":    "berhasil",
		"data_book": book,
	})
}

func CreateBookController(c echo.Context) error {
	book := model.Book{}
	c.Bind(&book)

	if err := config.DB.Save(&book).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new book",
		"book":    book,
	})
}

// delete user by id
func DeleteBookController(c echo.Context) error {
	// your solution here
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": "parameter salah",
		})
	}

	var book model.Book
	if err := config.DB.Where("id = ?", id).First(&book).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"status": "book not available",
		})
	}

	if err := config.DB.Delete(&book).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"status": "gagal hapus data"})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":    "berhasil menghapus data",
		"data_book": book,
	})

}

func UpdateBookController(c echo.Context) error {
	bady := new(model.Book)

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

	var book model.Book
	if err := config.DB.Where("id = ?", id).First(&book).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"status": "book not available",
		})
	}

	book.Judul = bady.Judul
	book.Penulis = bady.Penulis
	book.Penerbit = bady.Penerbit

	if err := config.DB.Save(&book).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"status": "gagal update data"})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "status berhasil update data",
		"status":  book,
	})
}
