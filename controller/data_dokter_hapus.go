package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"heroku.com/model"
)

func Data_dokter_hapus(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var dokter model.Dokter
	db.Where("id_user = ?", c.Param("id")).Find(&dokter)
	if dokter.Id_user == 0 {
		c.JSON(400, gin.H{
			"code":    400,
			"message": "Parameter id yang anda masukan salah.",
		})
		return
	}
	db.Delete(&dokter)
	c.JSON(200, gin.H{
		"code":    200,
		"data":    dokter,
		"message": "Data dokter berhasil sihapus.",
	})
}
