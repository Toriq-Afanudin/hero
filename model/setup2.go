package model

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
)

const (
	host     = "ec2-3-224-8-189.compute-1.amazonaws.com"
	user     = "mbakdsfqzwdlfw"
	password = "fff010e514d545f61dbd217951c2a3ef81fc2ef85f0349f015459d4f26b570e5"
	dbname   = "dfrk4g2r7718iv"
	post     = 5432
)

func SetupModels() *gorm.DB {
	var dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=require TimeZone=Asia/Jakarta", host, user, password, dbname, post)
	db, err := gorm.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected!")
	db.AutoMigrate(&Pasien{}, &Jadwal{}, &Dokter{}, &Rekam_medis{}, &Obat{}, &Ruangan{}, &Perawat{}, &User{})
	return db
}
