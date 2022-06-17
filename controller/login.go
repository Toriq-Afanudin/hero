package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"heroku.com/model"
)

type login struct {
	Email    string `form:"email" json:"email"`
	Password string `form:"password" json:"password"`
}

func Login(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var login login
	if err := c.ShouldBindJSON(&login); err != nil {
		c.JSON(400, gin.H{
			"status":  "error",
			"message": "input tidak dalam bentuk json",
		})
		return
	}
	var user model.User
	db.Where("email = ?", login.Email).Where("password = ?", login.Password).Find(&user)
	if login.Email == user.Email {
		c.JSON(200, gin.H{
			"status": "login berhasil",
			"email":  user.Email,
			"level":  user.Level,
		})
	} else {
		c.JSON(200, gin.H{
			"status": "email atau password salah",
		})
	}
}
