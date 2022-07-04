package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"heroku.com/model"
)

func Akun_hapus(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var akun model.User
	db.Where("id = ?", c.Param("id")).Find(&akun)
	if akun.Id == 0 {
		c.JSON(400, gin.H{
			"code":    400,
			"data":    "-",
			"message": "Parameter id yang anda masukan salah.",
		})
		return
	}
	db.Where("id = ?", c.Param("id")).Delete(&akun)
	c.JSON(200, gin.H{
		"code":    200,
		"data":    akun,
		"message": "Data user berhasil dihapus.",
	})
}
