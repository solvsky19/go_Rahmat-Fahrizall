package config

import (
	"fmt"
	"praktikum/lib/seeder"
	"praktikum/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	InitDB()
	InitialMigration()
}

func InitDB() {

	DB_Username := "root"
	DB_Password := "amek0086"
	DB_Port := "3306"
	DB_Host := "127.0.0.1"
	DB_Name := "crud_go"

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		DB_Username,
		DB_Password,
		DB_Host,
		DB_Port,
		DB_Name,
	)

	var err error
	DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{
		// Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err)
	}
	seeder.DBSeed(DB)

}

func InitialMigration() {
	DB.AutoMigrate(&models.Users{})
	DB.AutoMigrate(&models.Musics{})
	DB.AutoMigrate(&models.Tikets{})
	DB.AutoMigrate(&models.InformasiAcaras{})

}
