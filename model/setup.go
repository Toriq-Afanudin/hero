package model

import (
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func SetupModels() *gorm.DB {
	port := os.Getenv("PORT2")
	db, err := gorm.Open("mysql", "alterra:Udin@123@(hero2000.mysql.database.azure.com:%s)/capstone?charset=utf8mb4&parseTime=True&loc=Local", port)
	if err != nil {
		panic("gagal koneksi database")
	}
	db.AutoMigrate(&Pasien{}, &Jadwal{}, &Dokter{}, &Rekam_medis{}, &Obat{}, &Ruangan{}, &Perawat{}, &User{})
	return db
}
