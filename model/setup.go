package model

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
)

// const (
// 	host     = "hero2000.mysql.database.azure.com"
// 	database = "capstone"
// 	user     = "alterra"
// 	password = "Udin@123"
// )

func SetupModels() *gorm.DB {
	godotenv.Load()
	user := os.Getenv("USER")
	password := os.Getenv("PASSWORD")
	host := os.Getenv("SERVER")
	schema := os.Getenv("SCHEMA")
	var connectionString = fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?allowNativePasswords=true", user, password, host, schema)
	db, err := gorm.Open("mysql", connectionString)
	if err != nil {
		panic("gagal koneksi database")
	}
	db.AutoMigrate(&Pasien{}, &Jadwal{}, &Dokter{}, &Rekam_medis{}, &Obat{}, &Ruangan{}, &Perawat{}, &User{})
	return db
}
