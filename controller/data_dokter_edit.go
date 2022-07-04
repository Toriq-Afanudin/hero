package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"heroku.com/model"
)

type eDokter struct {
	Sip            string `json:"sip"`
	Nama           string `json:"nama"`
	Jenis_kelamin  string `json:"jenis_kelamin"`
	Spesialis      string `json:"spesialis"`
	Jadwal_praktek string `json:"jadwal_praktek"`
	Nomor_str      string `json:"nomor_str"`
	Nomor_telfon   string `json:"nomor_telfon"`
}

func Data_dokter_edit(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var eDokter eDokter
	if err := c.ShouldBindJSON(&eDokter); err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"data":    "-",
			"message": "Request harus dalam bentuk JSON.",
		})
		return
	}
	var dokter model.Dokter
	db.Where("id_user = ?", c.Param("id")).Find(&dokter)
	if dokter.Id_user == 0 {
		c.JSON(400, gin.H{
			"code":    400,
			"data":    "-",
			"message": "Parameter id yang anda masukan salah.",
		})
		return
	}
	if (eDokter.Spesialis == "umum") || (eDokter.Spesialis == "gigi") || (eDokter.Spesialis == "tht") || (eDokter.Spesialis == "kulit") {
	} else {
		c.JSON(400, gin.H{
			"code":    400,
			"data":    "-",
			"message": "Spesialis dokter yang dibutuhkan: umum, gigi, kulit, tht.",
		})
		return
	}
	db.Model(&dokter).Where("id_user = ?", c.Param("id")).Update("sip", eDokter.Sip)
	db.Model(&dokter).Where("id_user = ?", c.Param("id")).Update("nama_dokter", eDokter.Nama)
	db.Model(&dokter).Where("id_user = ?", c.Param("id")).Update("jenis_kelamin", eDokter.Jenis_kelamin)
	db.Model(&dokter).Where("id_user = ?", c.Param("id")).Update("spesialis", eDokter.Spesialis)
	db.Model(&dokter).Where("id_user = ?", c.Param("id")).Update("jadwal_praktek", eDokter.Jadwal_praktek)
	db.Model(&dokter).Where("id_user = ?", c.Param("id")).Update("nomor_str", eDokter.Nomor_str)
	db.Model(&dokter).Where("id_user = ?", c.Param("id")).Update("nomor_telfon", eDokter.Nomor_telfon)
	c.JSON(200, gin.H{
		"code":    200,
		"data":    eDokter,
		"message": "Data dokter berhasil dirubah.",
	})
}
