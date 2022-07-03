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
	var dokter model.Dokter
	db.Where("id_user = ?", c.Param("id")).Find(dokter)
	if dokter.Id_user == 0 {
		c.JSON(400, gin.H{
			"code":    400,
			"message": "Parameter id yang anda masukan salah.",
		})
		return
	}
	var eDokter eDokter
	if err := c.ShouldBindJSON(&eDokter); err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"message": "Request harus dalam bentuk JSON.",
		})
		return
	}
	db.Model(&dokter).Update("sip", eDokter.Sip)
	db.Model(&dokter).Update("nama_dokter", eDokter.Nama)
	db.Model(&dokter).Update("jenis_kelamin", eDokter.Jenis_kelamin)
	db.Model(&dokter).Update("spesialis", eDokter.Spesialis)
	db.Model(&dokter).Update("jadwal_praktek", eDokter.Jadwal_praktek)
	db.Model(&dokter).Update("nomor_str", eDokter.Nomor_str)
	db.Model(&dokter).Update("nomor_telfon", eDokter.Nomor_telfon)
	c.JSON(200, gin.H{
		"code": 200,
		"data": eDokter,
	})
}
