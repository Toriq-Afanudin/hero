package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"heroku.com/controller"
	"heroku.com/model"
)

func main() {
	r := gin.Default()

	db := model.SetupModels()
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	r.GET("/", controller.Utama)
	r.POST("/tambah_akun", controller.TambahAkun)

	godotenv.Load()
	ip := os.Getenv("IP")
	port := os.Getenv("PORT")
	address := fmt.Sprintf("%s:%s", ip, port)
	fmt.Println(address)

	r.Run(address)
}
