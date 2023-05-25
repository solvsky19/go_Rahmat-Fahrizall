package database

import (
	"praktikum/config"
	"praktikum/models"

	"gorm.io/gorm"
)

type CategoryRepository interface {
	GetAllCategory() (category []models.CategoryItems, err error)
	GetCategoryById(id int) (category *models.CategoryItems, err error)
	CreateCategory(category *models.CategoryItems) error
}

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *categoryRepository {
	return &categoryRepository{db}
}

func (c *categoryRepository) GetAllCategory() (category []models.CategoryItems, err error) {
	if err := config.DB.Preload("Items").Find(&category).Error; err != nil {
		return nil, err
	}

	return category, nil
}

func (c *categoryRepository) GetCategoryById(id int) (category *models.CategoryItems, err error) {
	if err := config.DB.Preload("Items").Where("id = ?", id).First(&category).Error; err != nil {
		return nil, err
	}

	return category, nil
}

func (c *categoryRepository) CreateCategory(category *models.CategoryItems) error {
	if err := config.DB.Create(&category).Error; err != nil {
		return err
	}

	return nil
}