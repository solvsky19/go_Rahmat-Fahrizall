package models

import (
	"time"

	"gorm.io/gorm"
)

type InformasiAcaras struct {
	ID        uint   `gorm:"primaryKey"`
	Lokasi    string `json:"Lokasi" form:"Lokasi"`
	Tanggal   string `json:"Tanggal" form:"Tanggal"`
	Waktu     string `json:"Waktu" form:"Waktu"`
	Harga     string `json:"Harga" form:"Harga"`
	DeletedAt gorm.DeletedAt
	CreatedAt time.Time
	UpdatedAt time.Time
}
