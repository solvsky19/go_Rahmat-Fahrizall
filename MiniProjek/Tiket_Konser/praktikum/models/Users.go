package models

import (
	"time"

	"gorm.io/gorm"
)

type Users struct {
	ID        int    `gorm:"primaryKey"`
	Name      string `json:"name" form:"name" gorm:"unique;not null"`
	Email     string `json:"email" form:"email"`
	Password  string `json:"password" form:"password"`
	Token     string `json:"token" form:"token"`
	DeletedAt gorm.DeletedAt
	CreatedAt time.Time
	UpdatedAt time.Time

	Role string `json:"role" form:"role"`
}
