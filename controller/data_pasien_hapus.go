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
	db.Where("id = ?", c.Param("id")).Delete(&pasien)
	var rekam model.Rekam_medis
	db.Where("id = ?", c.Param("id")).Find(&rekam)
	if rekam.Poli != "" {
		db.Where("id = ?", c.Param("id")).Delete(&rekam)
	}
	var rawat model.Rawat_jalan
	db.Where("id = ?", c.Param("id")).Find(&rawat)
	if rawat.Poli != "" {
		db.Where("id = ?", c.Param("id")).Delete(&rawat)
	}
	c.JSON(200, gin.H{
		"code": 200,
		"data": pasien,
	})
}
