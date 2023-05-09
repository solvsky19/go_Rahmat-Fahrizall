package models

import (
	"time"

	"gorm.io/gorm"
)

type Users struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `json:"name" form:"name"`
	Email     string `json:"email" form:"email"`
	Password  string `json:"password" form:"password"`
	Token     string `json:"token" form:"token"`
	DeletedAt gorm.DeletedAt
	CreatedAt time.Time
	UpdatedAt time.Time
}
