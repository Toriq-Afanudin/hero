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

func TambahAkun(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var b User
	if err := c.ShouldBindJSON(&b); err != nil {
		c.JSON(400, gin.H{
			"status":  "error",
			"message": "input tidak dalam bentuk json",
		})
		return
	}
	if b.Nickname == "" {
		c.JSON(400, gin.H{
			"status":  "error",
			"message": "data kosong",
		})
		return
	}
	create := model.User{
		Nickname: b.Nickname,
		Email:    b.Email,
		Password: b.Password,
		Level:    b.Level,
	}
	db.Create(&create)
	c.JSON(200, gin.H{
		"status": "akun berhasil di tambahkan",
		"akun":   b.Nickname,
	})
}
