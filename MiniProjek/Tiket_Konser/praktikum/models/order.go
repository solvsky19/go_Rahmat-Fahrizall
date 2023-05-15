package models

type Order struct {
	ID      uint `gorm:"primary_key" json:"id"`
	UserID  uint `json:"user_id"`
	TiketID uint `json:"tiket_id"`
	MusicID uint `json:"music_id"`
}
