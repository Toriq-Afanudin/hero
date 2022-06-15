package model

import (
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func SetupModels() *gorm.DB {
	database := os.Getenv("DATABASE")
	user := os.Getenv("USER")
	password := os.Getenv("PASSWORD")
	server := os.Getenv("SERVER")
	schema := os.Getenv("SCHEMA")
	port := os.Getenv("PORT2")
	db, err := gorm.Open("%s", "%s:%s@(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", database, user, password, server, port, schema)
	if err != nil {
		panic("gagal koneksi database")
	}
	db.AutoMigrate(&Pasien{}, &Jadwal{}, &Dokter{}, &Rekam_medis{}, &Obat{}, &Ruangan{}, &Perawat{}, &User{})
	return db
}
