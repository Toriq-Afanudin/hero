package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"heroku.com/model"
)

func Hapus_data_pasien(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var pasien model.Pasien
	db.Where("id = ?", c.Param("id")).Find(&pasien)
	if pasien.Nama == "" {
		c.JSON(400, gin.H{
			"status":  "Error",
			"message": "Parameter id tidak ditemukan.",
		})
		return
	}
	db.Delete(&pasien)
	c.JSON(200, gin.H{
		"code": 200,
		"data": pasien,
	})
}
