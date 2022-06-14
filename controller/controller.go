package controller

import (
	"github.com/gin-gonic/gin"
)

func Utama(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "selamat datang di aplikasi heroku",
	})
}
