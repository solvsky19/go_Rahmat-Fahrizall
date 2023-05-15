package database

import (
	"praktikum/config"
	"praktikum/models"
)

func CreateOrder(order models.Order) (models.Order, error) {
	err := config.DB.Create(&order).Error

	if err != nil {
		return models.Order{}, err
	}
	return order, nil
}

func GetOrderAll() ([]models.Order, error) {
	var order []models.Order

	if err := config.DB.Find(&order).Error; err != nil {
		return nil, err
	}
	return order, nil
}

func GetOrderByID(id int) (order models.Order, err error) {
	err = config.DB.Where("id = ?", id).Find(&order).Error
	if err != nil {
		return order, err
	}
	return order, nil
}

func CheckAllOrder(order models.Order) (models.Order, error) {
	err := config.DB.Where("user_id = ?", order.UserID).First(&order).Error
	if err != nil {
		return models.Order{}, err
	}
	return order, nil
}

func UpdateOrderByID(id uint, order *models.Order) (*models.Order, error) {
	var orders models.Order

	err := config.DB.First(&orders, id).Error
	if err != nil {
		return nil, err
	}

	orders.UserID = order.UserID
	orders.TiketID = order.TiketID
	orders.MusicID = order.MusicID

	err = config.DB.Save(&orders).Error
	if err != nil {
		return nil, err
	}

	return &orders, nil
}

func DeleteOrderByID(id uint) error {
	var orderss models.Order

	err := config.DB.First(&orderss, id).Error
	if err != nil {
		return err
	}

	err = config.DB.Delete(&orderss).Error
	if err != nil {
		return err
	}

	return nil
}
