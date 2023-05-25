package usecase

import (
	"praktikum/models"
	"praktikum/repository/database"
)

type CatergoryUsecase interface {
	GetAllCategory() (category []models.CategoryItems, err error)
	GetCategoryById(id int) (category *models.CategoryItems, err error)
	CreateCategory(category *models.CategoryItems) error
}

type categoryUsecase struct {
	categoryRepository database.CategoryRepository
}

func NewCategoryUsecase(categoryRepository database.CategoryRepository) *categoryUsecase {
	return &categoryUsecase{categoryRepository}
}

func (c *categoryUsecase) GetAllCategory() (category []models.CategoryItems, err error) {
	category, err = c.categoryRepository.GetAllCategory()
	if err != nil {
		return nil, err
	}

	return category, nil
}

func (c *categoryUsecase) GetCategoryById(id int) (category *models.CategoryItems, err error) {
	category, err = c.categoryRepository.GetCategoryById(id)
	if err != nil {
		return nil, err
	}

	return category, nil
}

func (c *categoryUsecase) CreateCategory(category *models.CategoryItems) error {
	err := c.categoryRepository.CreateCategory(category)
	if err != nil {
		return err
	}

	return nil
}