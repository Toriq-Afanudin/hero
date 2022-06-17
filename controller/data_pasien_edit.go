package controller

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"heroku.com/model"
)

type edit_data struct {
	Nik           string `json:"nik"`
	Nama          string `json:"nama"`
	Alamat        string `json:"alamat"`
	Jenis_kelamin string `json:"jenis_kelamin"`
	Nomer_telfon  string `json:"nomer_telfon"`
	Tempat_lahir  string `json:"tempat_lahir"`
	Tanggal_lahir string `json:"tanggal_lahir"`
}

func Edit_data_pasien(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	db := c.MustGet("db").(*gorm.DB)
	var user model.User
	db.Where("email = ?", claims["id"]).Where("level = ?", "admin").Find(&user)
	if claims["id"] == user.Email {
		c.JSON(400, gin.H{
			"status":  "gagal menampilkan data",
			"message": "yang berhak mengakses halaman ini hanya dokter atau perawat",
		})
		return
	}
	var edit edit_data
	if err := c.ShouldBindJSON(&edit); err != nil {
		c.JSON(400, gin.H{
			"status":  "error",
			"message": "input tidak dalam bentuk json",
		})
		return
	}
	var pasien model.Pasien
	db.Where("id = ?", c.Param("id")).Find(&pasien)
	if pasien.Nama == "" {
		c.JSON(400, gin.H{
			"status":  "error",
			"message": "parameter id yang anda masukan salah",
		})
		return
	}
	db.Model(&pasien).Update("nik", edit.Nik)
	db.Model(&pasien).Update("nama", edit.Nama)
	db.Model(&pasien).Update("alamat", edit.Alamat)
	db.Model(&pasien).Update("jenis_kelamin", edit.Jenis_kelamin)
	db.Model(&pasien).Update("no_hp", edit.Nomer_telfon)
	db.Model(&pasien).Update("tempat_lahir", edit.Tempat_lahir)
	db.Model(&pasien).Update("tanggal_lahir", edit.Tanggal_lahir)
	c.JSON(200, gin.H{
		"status": "berhasil mengedit data pasien",
		"data":   pasien,
		"userID": claims["id"],
	})
}
