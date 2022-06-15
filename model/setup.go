package model

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	// "github.com/joho/godotenv"
)

// const (
// 	host     = "hero2000.mysql.database.azure.com"
// 	database = "capstone"
// 	user     = "alterra"
// 	password = "Udin@123"
// )

func SetupModels() *gorm.DB {
	// godotenv.Load()
	user := os.Getenv("USER")
	password := os.Getenv("PASSWORD")
	host := os.Getenv("SERVER")
	schema := os.Getenv("SCHEMA")
	database := os.Getenv("DATABASE")
	port := os.Getenv("PORT")
	var connectionString = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?allowNativePasswords=true/sslmode=require", user, password, host, port, schema)
	var rdbms = fmt.Sprintf("%s", database)
	db, err := gorm.Open(rdbms, connectionString)
	if err != nil {
		panic("gagal koneksi database")
	}
	db.AutoMigrate(&Pasien{}, &Jadwal{}, &Dokter{}, &Rekam_medis{}, &Obat{}, &Ruangan{}, &Perawat{}, &User{})
	return db
}
