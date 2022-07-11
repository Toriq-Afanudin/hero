package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"heroku.com/model"
)

func Rawat_hapus(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var rawat model.Rawat_jalan
	db.Where("id = ?", c.Param("id")).Find(&rawat)
	if rawat.Id == 0 {
		c.JSON(400, gin.H{
			"code":    400,
			"data":    "-",
			"message": "Parameter id yang anda masukan salah.",
		})
		return
	}
	db.Where("id = ?", c.Param("id")).Delete(&rawat)
	c.JSON(200, gin.H{
		"code":    200,
		"data":    rawat,
		"message": "Data user berhasil dihapus.",
	})
}
