package controller

import (
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

func Akun_tambah(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var tambah User
	if err := c.ShouldBindJSON(&tambah); err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"message": "Request harus dalam bentuk JSON.",
		})
		return
	}
	if (tambah.Nickname == "") || (tambah.Email == "") || (tambah.Password == "") || (tambah.Level == "") {
		c.JSON(400, gin.H{
			"code":    400,
			"message": "Data tidak boleh ada yang kosong.",
		})
		return
	}
	var user model.User
	db.Where("nickname = ?", tambah.Nickname).Find(&user)
	if user.Nickname == tambah.Nickname {
		c.JSON(400, gin.H{
			"code":    400,
			"message": "Nickname sudah digunakan.",
		})
		return
	}
	db.Where("email = ?", tambah.Email).Find(&user)
	if user.Email == tambah.Email {
		c.JSON(400, gin.H{
			"code":    400,
			"message": "Email sudah digunakan.",
		})
		return
	}
	if len(tambah.Password) < 8 {
		c.JSON(400, gin.H{
			"code":    400,
			"message": "Password minimal terdiri dari 8 karakter.",
		})
		return
	}
	var pekerjaan bool
	if (tambah.Level == "perawat") || (tambah.Level == "dokter") || (tambah.Level == "admin") {
		pekerjaan = true
	} else {
		c.JSON(400, gin.H{
			"code":    400,
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
			message = "Dokter " + tambah.Nickname + " diharuskan melengkapi data sendiri"
		}
		if tambah.Level == "perawat" {
			new := model.Perawat{
				Id_user: mod.Id,
			}
			db.Create(&new)
			message = "Perawat " + tambah.Nickname + " diharuskan melengkapi data sendiri"
		}
		if tambah.Level == "admin" {
			message = tambah.Nickname + " resmi menjadi admin baru"
		}
		c.JSON(200, gin.H{
			"code":    200,
			"data":    tambah,
			"message": message,
		})
	}
}
