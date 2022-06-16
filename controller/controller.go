package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"heroku.com/model"
)

type login struct {
	Email    string `form:"email"`
	Password string `form:"password"`
}

func Utama(c *gin.Context) {
	c.JSON(200, gin.H{
		"data": "selamat datang di aplikasi hospital",
	})
}

func GetUser(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var user []model.User
	db.Find(&user)
	c.JSON(200, gin.H{
		"data": user,
	})
}

func Login(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var l login
	var user model.User
	db.Where("email = ?", l.Email).Where("password = ?", l.Password).Find(&user)
	if l.Email != user.Email {
		c.JSON(400, gin.H{
			"status": "email atau password salah",
		})
	}
	if user.Level == 1 {
		c.JSON(400, gin.H{
			"level": "dokter/perawat",
			"email": l.Email,
		})
	}
	if user.Level == 2 {
		c.JSON(400, gin.H{
			"level": "admin",
			"email": l.Email,
		})
	}
}
