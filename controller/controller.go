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
	var l login
	if (l.Email == "admin@gmail.com") && (l.Password == "admin123") {
		c.JSON(200, gin.H{
			"level": "admin",
			"email": l.Email,
		})
	} else if (l.Email == "dokter@gmail.com") && (l.Password == "dokter123") {
		c.JSON(200, gin.H{
			"level": "dokter",
			"email": l.Email,
		})
	} else {
		c.JSON(400, gin.H{
			"status": "email atau password salah",
		})
	}
}
