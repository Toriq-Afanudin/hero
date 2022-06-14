package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func SetupModels() *gorm.DB {
	db, err := gorm.Open("mysql", "toriq:Ayu!1999@(toriq1999.mysql.database.azure.com:3306)/capstone?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic("gagal koneksi database")
	}
	db.AutoMigrate(&User{})
	return db
}
