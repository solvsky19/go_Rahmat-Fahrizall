package models

type CategoryItems struct {
	ID           uint   `json:"id" form:"id" gorm:"primarykey"`
	NameCategory string `json:"name" form:"name"`
	Items        []Item `gorm:"foreignKey:CategoryID"`
}