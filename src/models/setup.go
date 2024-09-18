package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectData() {
	dsn := "root:@tcp(localhost:3306)/golang_api?charset=utf8mb4&parseTime=True&loc=Local"
	database, err := gorm.Open(mysql.Open(dsn))

	if err != nil {
		panic("Gagal menghubungkan ke database: " + err.Error())
	}

	database.AutoMigrate(&Product{})

	DB = database
}
