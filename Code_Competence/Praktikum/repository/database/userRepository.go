package database

import (
	"praktikum/config"
	"praktikum/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetUser() (user []models.User, err error)
	GetUserByUsername(username string) (user *models.User, err error)
	GetUserById(id int) (user *models.User, err error)
	Createuser(user *models.User) error
	UpdateUser(user *models.User) error
	DeleteUser(user *models.User) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (u *userRepository) GetUser() (user []models.User, err error) {
	err = config.DB.Find(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *userRepository) GetUserByUsername(username string) (user *models.User, err error) {
	err = config.DB.Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *userRepository) GetUserById(id int) (user *models.User, err error) {
	err = config.DB.Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *userRepository) Createuser(user *models.User) error {
	err := config.DB.Create(&user).Error
	if err != nil {
		return err
	}
	return nil
}

func (u *userRepository) UpdateUser(user *models.User) error {
	err := config.DB.Updates(&user).Error
	if err != nil {
		return err
	}
	return nil
}

func (u *userRepository) DeleteUser(user *models.User) error {
	err := config.DB.Delete(&user).Error
	if err != nil {
		return err
	}
	return nil
}