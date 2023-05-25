package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `json:"username" form:"username" gorm:"unique"`
	Password string `json:"password" form:"password"`
	Token    string `json:"-" gorm:"-"`
}