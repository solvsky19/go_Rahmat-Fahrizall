package database

import (
	"praktikum/config"
	"praktikum/models"

	"gorm.io/gorm"
)

type ItemRepository interface {
	GetAllItem() (item []models.Item, err error)
	GetItemById(id string) (item *models.Item, err error)
	CreateItem(item *models.Item) error
	UpdateItemById(id string, item *models.Item) error
	DeleteItemById(id string, item *models.Item) error
	GetItemByCategoryId(id int) (item []models.Item, err error)
	GetItemByName(name string) (item []models.Item, err error)
}

type itemRepository struct {
	db *gorm.DB
}

func NewItemRepository(db *gorm.DB) *itemRepository {
	return &itemRepository{db}
}

func (i *itemRepository) GetAllItem() (item []models.Item, err error) {
	if err := config.DB.Find(&item).Error; err != nil {
		return nil, err
	}

	return item, nil
}

func (i *itemRepository) GetItemById(id string) (item *models.Item, err error) {
	if err := config.DB.Where("id = ?", id).First(&item).Error; err != nil {
		return nil, err
	}

	return item, nil
}

func (i *itemRepository) GetItemByName(name string) (item []models.Item, err error) {
	if err := config.DB.Where("name like ?", "%"+name+"%").Find(&item).Error; err != nil {
		return nil, err
	}

	return item, nil
}

func (i *itemRepository) GetItemByCategoryId(id int) (item []models.Item, err error) {
	if err := config.DB.Where("category_id = ?", id).Find(&item).Error; err != nil {
		return nil, err
	}

	return item, nil
}

func (i *itemRepository) CreateItem(item *models.Item) error {
	if err := config.DB.Create(&item).Error; err != nil {
		return err
	}

	return nil
}

func (i *itemRepository) UpdateItemById(id string, item *models.Item) error {
	if err := config.DB.Where("id = ?", id).Updates(&item).Error; err != nil {
		return err
	}

	return nil
}

func (i *itemRepository) DeleteItemById(id string, item *models.Item) error {
	if err := config.DB.Where("id = ?", id).Delete(&item).Error; err != nil {
		return err
	}

	return nil
}
