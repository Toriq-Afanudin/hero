package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"heroku.com/model"
)

type login struct {
	Email    string `form:"email"`
	Password string `form:"password"`
}

func Login(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var l login
	if err := c.ShouldBindJSON(&l); err != nil {
		c.JSON(400, gin.H{
			"status":  "error",
			"message": "input tidak dalam bentuk json",
		})
		return
	}
	var u model.User
	db.Where("email = ?", l.Email).Where("password = ?", l.Password).Find(&u)
	if l.Email == u.Email {
		c.JSON(200, gin.H{
			"status": "login berhasil",
			"email":  u.Email,
			"level":  u.Level,
		})
	} else {
		c.JSON(200, gin.H{
			"status": "email atau password salah",
		})
	}
}
