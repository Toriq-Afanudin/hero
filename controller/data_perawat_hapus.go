package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"heroku.com/model"
)

func Data_perawat_hapus(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var perawat model.Perawat
	db.Where("id_user = ?", c.Param("id")).Find(&perawat)
	if perawat.Id_user == 0 {
		c.JSON(400, gin.H{
			"code":    400,
			"message": "Parameter id yang anda masukan salah.",
		})
		return
	}
	db.Delete(&perawat)
	c.JSON(200, gin.H{
		"code":    200,
		"data":    perawat,
		"message": "Data perawat berhasil sihapus.",
	})
}
