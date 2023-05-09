package models

import (
	"time"

	"gorm.io/gorm"
)

type Tikets struct {
	ID         uint   `gorm:"primaryKey"`
	JenisTiket string `json:"JenisTiket" form:"JenisTiket"`
	Penyanyi   string `json:"Penyanyi" form:"Penyanyi"`
	Waktu      string `json:"Waktu" form:"Waktu"`
	DeletedAt  gorm.DeletedAt
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
