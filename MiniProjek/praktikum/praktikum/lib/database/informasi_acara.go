package database

import (
	"praktikum/config"
	"praktikum/models"
)

func GetAllInformasiAcaras() ([]models.InformasiAcaras, error) {
	var users []models.InformasiAcaras

	err := config.DB.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func GetInformasiAcaraById(id uint) (*models.InformasiAcaras, error) {
	var user models.InformasiAcaras

	err := config.DB.First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func CreateInformasiAcaras(user *models.InformasiAcaras) (*models.InformasiAcaras, error) {

	err := config.DB.Create(&user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}
func UpdateInformasiAcaraById(id uint, informasiacara *models.InformasiAcaras) (*models.InformasiAcaras, error) {

	var record models.InformasiAcaras

	err := config.DB.First(&record, id).Error
	if err != nil {
		return nil, err
	}

	record.Lokasi = informasiacara.Lokasi
	record.Tanggal = informasiacara.Tanggal
	record.Waktu = informasiacara.Waktu
	record.Harga = informasiacara.Harga

	err = config.DB.Save(&record).Error
	if err != nil {
		return nil, err
	}

	return &record, nil
}

func DeleteInformasiAcaraById(id uint) error {
	var user models.InformasiAcaras

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
