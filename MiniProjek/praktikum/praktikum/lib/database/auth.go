package database

import (
	"fmt"
	"praktikum/config"
	"praktikum/middlewares"
	"praktikum/models"
)

func LoginUsers(user *models.Users) (interface{}, error) {

	var err error
	if err = config.DB.Where("email= ? AND password = ?",
		user.Email, user.Password).First(user).Error; err != nil {
		return nil, fmt.Errorf("user not found")

	}

	token, err := middlewares.CreateToken(int(user.ID))
	if err != nil {
		return "", fmt.Errorf("failed to create token: %v", err)
	}

	return token, nil
}
