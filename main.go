package main

import (
	"os"

	"fmt"

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

	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, X-Max")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	})

	r.GET("/", controller.Utama)

	r.POST("/login", controller.Login)

	admin := r.Group("/admin")
	admin.POST("/akun_tambah", controller.Akun_tambah)
	admin.GET("/akun_tampil", controller.Akun_tampil)
	admin.DELETE("/akun_hapus/:id", controller.Akun_hapus)

	admin.POST("/data_pasien_tambah", controller.Tambah_data_pasien)
	admin.GET("/data_pasien", controller.DataPasien)
	admin.PUT("/data_pasien_edit/:id", controller.Edit_data_pasien)
	admin.DELETE("/data_pasien_hapus/:id", controller.Hapus_data_pasien)

	admin.GET("/rawat_jalan_lihat", controller.Rawat_jalan_lihat)
	admin.PUT("/rawat_jalan_edit/:id", controller.Rawat_jalan_edit)

	admin.GET("/data_dokter_lihat", controller.Data_dokter_lihat)
	admin.PUT("/data_dokter_edit/:id", controller.Data_dokter_edit)
	admin.DELETE("/data_dokter_hapus/:id", controller.Data_dokter_hapus)

	admin.GET("/data_perawat_lihat", controller.Data_perawat_lihat)
	admin.PUT("/data_perawat_edit/:id", controller.Data_perawat_edit)
	admin.DELETE("/data_perawat_hapus/:id", controller.Data_perawat_hapus)

	dokter := r.Group("/dokter")
	dokter.PUT("/akun_dokter_update/:id", controller.Edit_akun_dokter_by_id)
	dokter.GET("/akun_dokter_lihat", controller.Lihat_akun_dokter)

	r.DELETE("/rekam/:id", controller.Rekam_hapus)
	r.DELETE("/rawat/:id", controller.Rawat_hapus)

	godotenv.Load()
	port := os.Getenv("PORT")
	var dns = fmt.Sprintf("0.0.0.0:%s", port)

	r.Run(dns)
}
