package controller

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"heroku.com/model"
)

type User struct {
	Nickname string `json:"nickname"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Level    string `json:"level"`
}

func TambahAkun(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	db := c.MustGet("db").(*gorm.DB)
	var use model.User
	db.Where("email = ?", claims["id"]).Where("level = ?", "admin").Find(&use)
	if claims["id"] == use.Email {
		c.JSON(400, gin.H{
			"status":  "Error",
			"message": "Halaman ini hanya bisa diakses oleh dokter atau perawat.",
		})
		return
	}
	var tambah User
	if err := c.ShouldBindJSON(&tambah); err != nil {
		c.JSON(400, gin.H{
			"status":  "Error",
			"message": "Request harus dalam bentuk JSON.",
		})
		return
	}
	if (tambah.Nickname == "") || (tambah.Email == "") || (tambah.Password == "") || (tambah.Level == "") {
		c.JSON(400, gin.H{
			"status":  "Error",
			"message": "Data tidak boleh ada yang kosong.",
		})
		return
	}
	var user model.User
	db.Where("nickname = ?", tambah.Nickname).Find(&user)
	if user.Nickname == tambah.Nickname {
		c.JSON(400, gin.H{
			"status":  "Error",
			"message": "Nickname sudah digunakan.",
		})
		return
	}
	db.Where("email = ?", tambah.Email).Find(&user)
	if user.Email == tambah.Email {
		c.JSON(400, gin.H{
			"status":  "Error",
			"message": "Email sudah digunakan.",
		})
		return
	}
	if len(tambah.Password) < 8 {
		c.JSON(400, gin.H{
			"status":  "Error",
			"message": "Password minimal terdiri dari 8 karakter.",
		})
		return
	}
	var pekerjaan bool
	if (tambah.Level == "perawat") || (tambah.Level == "dokter") || (tambah.Level == "admin") {
		pekerjaan = true
	} else {
		c.JSON(400, gin.H{
			"status":  "Error",
			"message": "Pilih jenis pekerjaan sebagai admin, dokter, atau perawat.",
		})
		return
	}
	create := model.User{
		Nickname: tambah.Nickname,
		Email:    tambah.Email,
		Password: tambah.Password,
		Level:    tambah.Level,
	}
	if pekerjaan {
		db.Create(&create)
		var mod model.User
		db.Where("nickname = ?", tambah.Nickname).Find(&mod)
		var message string
		if tambah.Level == "dokter" {
			new := model.Dokter{
				Id_user: mod.Id,
			}
			db.Create(&new)
			message = "Lengkapi data dokter " + tambah.Nickname
		}
		if tambah.Level == "perawat" {
			new := model.Perawat{
				Id_user: mod.Id,
			}
			db.Create(&new)
			message = "Lengkapi data perawat " + tambah.Nickname
		}
		c.JSON(200, gin.H{
			"status":  "Berhasil",
			"data":    tambah.Nickname,
			"user":    claims["id"],
			"message": message,
		})
	}
}
