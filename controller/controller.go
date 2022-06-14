package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"heroku.com/model"
)

func Utama(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "selamat datang di aplikasi heroku",
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
