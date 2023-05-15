//login untuk admin

package facker

import (
	"praktikum/constants"
	"praktikum/models"

	"gorm.io/gorm"
)

func UserFaker(db *gorm.DB) *models.Users {
	return &models.Users{
		Name:     "admin",
		Email:    "admin@gmail.com",
		Password: "admin",
		Role:     constants.ADMIN,
	}
}
