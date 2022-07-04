package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"heroku.com/model"
)

type perawat struct {
	Sip           string `json:"sip"`
	Nama          string `json:"nama"`
	Jenis_kelamin string `json:"jenis_kelamin"`
	Bagian_kerja  string `json:"bagian_kerja"`
	Jadwal_kerja  string `json:"jadwal_kerja"`
	Jabatan       string `json:"jabatan"`
	Nomor_telfon  string `json:"nomor_telfon"`
	Nomor_str     string `json:"nomor_str"`
}

func Data_perawat_edit(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var mPerawat model.Perawat
	db.Where("id_user = ?", c.Param("id")).Find(&mPerawat)
	if mPerawat.Id_user == 0 {
		c.JSON(400, gin.H{
			"code":    400,
			"data":    "-",
			"message": "Parameter id tidak ditemukan.",
		})
		return
	}
	var perawat perawat
	if err := c.ShouldBindJSON(&perawat); err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"data":    "-",
			"message": "Request harus dalam bentuk JSON.",
		})
		return
	}
	if perawat.Bagian_kerja == "umum" {
	} else {
		c.JSON(400, gin.H{
			"code":    400,
			"data":    "-",
			"message": "Bagian kerja yang dibutuhkan: umum, gigi, kulit, tht.",
		})
		return
	}
	db.Model(&mPerawat).Where("id_user = ?", c.Param("id")).Update("sip", perawat.Sip)
	db.Model(&mPerawat).Where("id_user = ?", c.Param("id")).Update("nama_perawat", perawat.Nama)
	db.Model(&mPerawat).Where("id_user = ?", c.Param("id")).Update("jenis_kelamin", perawat.Jenis_kelamin)
	db.Model(&mPerawat).Where("id_user = ?", c.Param("id")).Update("bagian_kerja", perawat.Bagian_kerja)
	db.Model(&mPerawat).Where("id_user = ?", c.Param("id")).Update("jadwal_kerja", perawat.Jadwal_kerja)
	db.Model(&mPerawat).Where("id_user = ?", c.Param("id")).Update("jabatan", perawat.Jabatan)
	db.Model(&mPerawat).Where("id_user = ?", c.Param("id")).Update("nomor_telfon", perawat.Nomor_telfon)
	db.Model(&mPerawat).Where("id_user = ?", c.Param("id")).Update("nomor_str", perawat.Nomor_str)
	c.JSON(200, gin.H{
		"code":    200,
		"data":    perawat,
		"message": "Data perawat berhasil dirubah",
	})
}
