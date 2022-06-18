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
			"status":  "gagal menampilkan data",
			"message": "yang berhak mengakses halaman ini hanya dokter atau perawat",
		})
		return
	}
	var tambah User
	if err := c.ShouldBindJSON(&tambah); err != nil {
		c.JSON(400, gin.H{
			"status":  "error",
			"message": "input tidak dalam bentuk json",
		})
		return
	}
	if (tambah.Nickname == "") || (tambah.Email == "") || (tambah.Password == "") || (tambah.Level == "") {
		c.JSON(400, gin.H{
			"status":  "error",
			"message": "ada data yang belum di isi",
		})
		return
	}
	var user model.User
	db.Where("nickname = ?", tambah.Nickname).Find(&user)
	if user.Nickname == tambah.Nickname {
		c.JSON(400, gin.H{
			"status":  "error",
			"message": "nickname sudah digunakan",
		})
		return
	}
	db.Where("email = ?", tambah.Email).Find(&user)
	if user.Email == tambah.Email {
		c.JSON(400, gin.H{
			"status":  "error",
			"message": "email sudah digunakan",
		})
		return
	}
	if len(tambah.Password) < 8 {
		c.JSON(400, gin.H{
			"status":  "error",
			"message": "password minimal 8 karakter",
		})
		return
	}
	var pekerjaan bool
	if (tambah.Level == "perawat") || (tambah.Level == "dokter") || (tambah.Level == "admin") {
		pekerjaan = true
	} else {
		c.JSON(400, gin.H{
			"status":  "error",
			"message": "jenis pekerjaan harus di isi dengan perawat atau dokter atau admin",
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
			message = "berhasil menambahkan dokter baru, dengan akun " + tambah.Nickname + ", silakan untuk melengkapi data dokter"
		}
		if tambah.Level == "perawat" {
			new := model.Perawat{
				Id_user: mod.Id,
			}
			db.Create(&new)
			message = "berhasil menambahkan perawat baru, dengan akun " + tambah.Nickname + ", silakan untuk melengkapi data perawat"
		}
		c.JSON(200, gin.H{
			"status":   "akun berhasil di tambahkan",
			"nickname": tambah.Nickname,
			"email":    tambah.Email,
			"userID":   claims["id"],
			"message":  message,
		})
	}
}
