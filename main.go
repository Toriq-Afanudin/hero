package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"

	"heroku.com/controller"
	"heroku.com/model"
)

func main() {
	r := gin.Default()

	// db := model.SetupModels()
	// r.Use(func(c *gin.Context) {
	// 	c.Set("db", db)
	// 	c.Next()
	// })

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
	db.AutoMigrate(&model.Pasien{}, &model.Jadwal{}, &model.Dokter{}, &model.Rekam_medis{}, &model.Obat{}, &model.Ruangan{}, &model.Perawat{}, &model.User{})

	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	r.GET("/user", controller.GetUser)
	r.GET("/", controller.Utama)
	r.GET("/login", controller.Login)

	godotenv.Load()
	ip := os.Getenv("IP")
	port := os.Getenv("PORT")
	address := fmt.Sprintf("%s:%s", ip, port)
	fmt.Println(address)
	r.Run(address)
}
