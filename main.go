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
	port := os.Getenv("PORT")
	var dns = fmt.Sprintf("0.0.0.0:%s", port)

	r.Run(dns)
}
