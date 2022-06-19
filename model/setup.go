package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// func SetupModels() *gorm.DB {
// 	db, err := gorm.Open("mysql", "alterra:Udin@123@tcp(hero2000.mysql.database.azure.com:3306)/capstone?charset=utf8mb4&parseTime=True&loc=Local")
// 	if err != nil {
// 		panic("gagal koneksi database")
// 	}
// 	db.AutoMigrate(&Pasien{}, &Jadwal{}, &Dokter{}, &Rekam_medis{}, &Obat{}, &Ruangan{}, &Perawat{}, &User{})
// 	return db
// }

func SetupModels() *gorm.DB {
	db, err := gorm.Open("postgres", "host=ec2-3-224-8-189.compute-1.amazonaws.com user=mbakdsfqzwdlfw password=fff010e514d545f61dbd217951c2a3ef81fc2ef85f0349f015459d4f26b570e5 dbname=dfrk4g2r7718iv port=5432 sslmode=disable TimeZone=Asia/Shanghai")
	if err != nil {
		panic("gagal koneksi database")
	}
	db.AutoMigrate(&Pasien{}, &Jadwal{}, &Dokter{}, &Rekam_medis{}, &Obat{}, &Ruangan{}, &Perawat{}, &User{})
	return db
}
