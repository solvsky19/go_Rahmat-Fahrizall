package database

import (
	"praktikum/config"
	"praktikum/models"
)

func GetAllTikets() ([]models.Tikets, error) {
	var users []models.Tikets

	err := config.DB.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func GetTiketById(id uint) (*models.Tikets, error) {
	var user models.Tikets

	err := config.DB.First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func CreateTikets(user *models.Tikets) (*models.Tikets, error) {

	err := config.DB.Create(&user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}
func UpdateTiketById(id uint, tiket *models.Tikets) (*models.Tikets, error) {

	var record models.Tikets

	err := config.DB.First(&record, id).Error
	if err != nil {
		return nil, err
	}

	record.JenisTiket = tiket.JenisTiket
	record.Penyanyi = tiket.Penyanyi
	record.Waktu = tiket.Waktu
	record.Harga = tiket.Harga

	err = config.DB.Save(&record).Error
	if err != nil {
		return nil, err
	}

	return &record, nil
}

func DeleteTiketById(id uint) error {
	var user models.Tikets

	err := config.DB.First(&user, id).Error
	if err != nil {
		return err
	}

	err = config.DB.Delete(&user).Error
	if err != nil {
		return err
	}

	return nil
}

