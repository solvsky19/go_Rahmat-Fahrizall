package database

import (
	"praktikum/config"
	"praktikum/models"
)

func GetAllMusics() ([]models.Musics, error) {
	var musics []models.Musics

	err := config.DB.Find(&musics).Error
	if err != nil {
		return nil, err
	}
	return musics, nil
}

func GetMusicByID(id uint) (*models.Musics, error) {
	var music models.Musics

	err := config.DB.First(&music, id).Error
	if err != nil {
		return nil, err
	}
	return &music, nil
}

func CreateMusic(music *models.Musics) (*models.Musics, error) {
	err := config.DB.Create(&music).Error
	if err != nil {
		return nil, err
	}

	return music, nil
}

func UpdateMusicByID(id uint, music *models.Musics) (*models.Musics, error) {
	var record models.Musics

	err := config.DB.First(&record, id).Error
	if err != nil {
		return nil, err
	}

	record.Judul = music.Judul
	record.Penyanyi = music.Penyanyi
	record.Pencipta = music.Pencipta

	err = config.DB.Save(&record).Error
	if err != nil {
		return nil, err
	}

	return &record, nil
}

func DeleteMusicByID(id uint) error {
	var music models.Musics

	err := config.DB.First(&music, id).Error
	if err != nil {
		return err
	}

	err = config.DB.Delete(&music).Error
	if err != nil {
		return err
	}

	return nil
}
