package controller

import (
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func Utama(c *gin.Context) {
	c.JSON(200, gin.H{
		"data": "selamat datang di aplikasi hospital",
	})
}
