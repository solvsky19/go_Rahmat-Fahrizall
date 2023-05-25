package models

import "github.com/google/uuid"

type MYUUID = uuid.UUID

type Item struct {
	ID          MYUUID `json:"id" form:"id" gorm:"primarykey"`
	Name        string `json:"name" form:"name"`
	Description string `json:"description" form:"description"`
	Stock       uint   `json:"stock" form:"stock"`
	Price       uint   `json:"price" form:"price"`
	CategoryID  uint   `json:"category_id" form:"category_id"`
}
