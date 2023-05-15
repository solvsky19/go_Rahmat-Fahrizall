package models

import (
	"time"

	"gorm.io/gorm"
)

type Musics struct {
	ID        uint   `gorm:"primaryKey"`
	Judul     string `json:"Judul" form:"Judul"`
	Penyanyi  string `json:"Penyanyi" form:"Penyanyi"`
	Pencipta  string `json:"Pencipta" form:"Pencipta"`
	DeletedAt gorm.DeletedAt
	CreatedAt time.Time
	UpdatedAt time.Time
}
