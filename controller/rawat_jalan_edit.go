package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"heroku.com/model"
)

type edit struct {
	Jadwal_rawat_jalan string `json:"jadwal_rawat_jalan"`
}

type dokter_edit struct {
	Keterangan string `json:"keterangan"`
	Proses     bool   `json:"proses"`
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
			"code":    400,
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
	var dok_edit dokter_edit
	if err := c.ShouldBindJSON(&dok_edit); err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"message": "Request harus dalam bentuk JSON.",
		})
		return
	}
	var ra_jalan model.Rawat_jalan
	db.Model(&ra_jalan).Where("id_pasien = ?", c.Param("id_pasien")).Update("keterangan", dok_edit.Keterangan)
	db.Model(&ra_jalan).Where("id_pasien = ?", c.Param("id_pasien")).Update("bool", dok_edit.Proses)
	c.JSON(200, gin.H{
		"code":       200,
		"message":    "Berhasil merubah keterangan dan proses.",
		"keterangan": dok_edit.Keterangan,
		"proses":     dok_edit.Proses,
	})
}
