package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type akun struct {
	Id            string `json:"id"`
	Sip           string `json:"sip"`
	Nama          string `json:"nama"`
	Jenis_kelamin string `json:"jenis_kelamin"`
	Email         string `json:"email"`
	Password      string `json:"password"`
	Level         string `json:"perawat"`
}

func List_account(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var akun []akun
	db.Raw("SELECT id, email, password, level FROM capstone.users;").Scan(&akun)
	c.JSON(200, gin.H{
		"status": "berhasil menampilkan data pasien",
		"akun":   akun,
	})
}
