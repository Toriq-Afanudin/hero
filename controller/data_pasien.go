package controller

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"heroku.com/model"
)

type pasien struct {
	Id             int    `form:"id"`
	Nik            string `form:"nik"`
	Nama           string `form:"nama"`
	Alamat         string `form:"alamat"`
	Jenis_kelamin  string `form:"jenis_kelamin"`
	Jenis_penyakit interface{}
	No_hp          string `form:"no_hp"`
	Tempat_lahir   string `form:"tempat_lahir"`
	Tanggal_lahir  string `form:"tanggal_lahir"`
}

type penyakit struct {
	Pemeriksaan string `form:"pemeriksaan"`
}

func DataPasien(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	db := c.MustGet("db").(*gorm.DB)
	var user model.User
	db.Where("email = ?", claims["id"]).Where("level = ?", "admin").Find(&user)
	if claims["id"] == user.Email {
		c.JSON(400, gin.H{
			"status":  "Error",
			"message": "Halaman ini hanya bisa diakses oleh dokter atau perawat.",
		})
		return
	}
	var d []pasien
	db.Raw("SELECT * FROM capstone.pasiens;").Scan(&d)
	var p penyakit
	for i := 0; i < len(d); i++ {
		db.Raw("SELECT pemeriksaan FROM capstone.rekam_medis WHERE id=?", d[i].Id).Scan(&p)
		db.Model(&d[i]).Update("Jenis_penyakit", p.Pemeriksaan)
	}
	c.JSON(200, gin.H{
		"status": "Berhasil",
		"data":   d,
		"user":   claims["id"],
	})
}
