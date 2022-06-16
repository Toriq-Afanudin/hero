package model

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
)

func SetupModels() *gorm.DB {
	godotenv.Load()
	user := os.Getenv("USER")
	password := os.Getenv("PASSWORD")
	host := os.Getenv("SERVER")
	schema := os.Getenv("SCHEMA")
	var connectionString = fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, schema)
	db, err := gorm.Open("mysql", connectionString)
	if err != nil {
		panic("gagal koneksi database")
	}
	db.AutoMigrate(&Pasien{}, &Jadwal{}, &Dokter{}, &Rekam_medis{}, &Obat{}, &Ruangan{}, &Perawat{}, &User{})
	return db
}
