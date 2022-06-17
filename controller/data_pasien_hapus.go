package controller

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"heroku.com/model"
)

func Hapus_data_pasien(c *gin.Context) {
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
	var pasien model.Pasien
	db.Where("id = ?", c.Param("id")).Find(&pasien)
	if pasien.Nama == "" {
		c.JSON(400, gin.H{
			"status":  "error",
			"message": "parameter id yang anda masukan salah",
		})
		return
	}
	db.Delete(&pasien)
	c.JSON(200, gin.H{
		"status": "berhasil menghapus data pasien",
		"data":   pasien,
		"userID": claims["id"],
	})
}
