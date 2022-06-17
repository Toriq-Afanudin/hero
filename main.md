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
	r.GET("/login", controller.Login)
	r.GET("/data_pasien", controller.DataPasien)
	r.POST("/tambah_data_pasien", controller.Tambah_data_pasien)
	r.GET("/akun_user", controller.List_account)

	godotenv.Load()
	port := os.Getenv("PORT")
	var dns = fmt.Sprintf("127.0.0.1:%s", port)

	r.Run(dns)
}
