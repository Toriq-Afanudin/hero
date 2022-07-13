package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"heroku.com/model"
)

type edit struct {
	Jadwal_rawat_jalan string `json:"jadwal_rawat_jalan"`
}

func Rawat_jalan_edit(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var rJalan model.Rawat_jalan
	db.Where("id_pasien = ?", c.Param("id")).Find(&rJalan)
	if rJalan.Id == 0 {
		c.JSON(400, gin.H{
			"code":    400,
			"message": "Parameter id yang anda masukan salah",
		})
		return
	}
	var e edit
	if err := c.ShouldBindJSON(&e); err != nil {
		c.JSON(400, gin.H{
			"status":  "Error",
			"message": "Request harus dalam bentuk JSON.",
		})
		return
	}
	db.Model(&rJalan).Update("jadwal_rawat_jalan", e.Jadwal_rawat_jalan)
	c.JSON(200, gin.H{
		"code":    200,
		"message": "Jadwal berhasil dirubah",
	})
}

func Rawat_jalan_ubah_proses(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var rJalan model.Rawat_jalan
	db.Where("id_pasien = ?", c.Param("id_pasien")).Find(&rJalan)
	if rJalan.Id_pasien == 0 {
		c.JSON(400, gin.H{
			"code":    400,
			"message": "Parameter id_pasien salah atau tidak ditemukan",
		})
		return
	}
	if !rJalan.Bool {
		db.Model(&rJalan).Where("id_pasien = ?", c.Param("id_pasien")).Update("bool", true)
		c.JSON(200, gin.H{
			"code":    200,
			"message": "Proses = true",
		})
		return

	}
	if rJalan.Bool {
		db.Model(&rJalan).Where("id_pasien = ?", c.Param("id_pasien")).Update("bool", false)
		c.JSON(200, gin.H{
			"code":    200,
			"message": "Proses = false",
		})
		return
	}
}
