package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"heroku.com/model"
)

type dokter struct {
	Sip            string `json:"sip"`
	Nama_dokter    string `json:"nama_dokter"`
	Spesialis      string `json:"spesialis"`
	Jenis_kelamin  string `json:"jenis_kelamin"`
	Jadwal_praktek string `json:"jadwal_praktek"`
	Nomor_str      string `json:"nomor_str"`
	Nik            string `json:"nik"`
	Alamat         string `json:"alamat"`
}

func Edit_akun_dokter_by_id(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var dokter dokter
	if err := c.ShouldBindJSON(&dokter); err != nil {
		c.JSON(400, gin.H{
			"status":  "Error",
			"message": "Request harus dalam bentuk JSON.",
		})
		return
	}
	var mDokter model.Dokter
	db.Where("id_user = ?", c.Param("id")).Find(&mDokter)
	if mDokter.Id_user == 0 {
		c.JSON(400, gin.H{
			"code":    400,
			"message": "Parameter salah.",
		})
		return
	}
	db.Model(&mDokter).Update("sip", dokter.Sip)
	db.Model(&mDokter).Update("nama_dokter", dokter.Nama_dokter)
	db.Model(&mDokter).Update("spesialis", dokter.Spesialis)
	db.Model(&mDokter).Update("jenis_kelamin", dokter.Jenis_kelamin)
	db.Model(&mDokter).Update("jadwal_praktek", dokter.Jadwal_praktek)
	db.Model(&mDokter).Update("nomor_str", dokter.Nomor_str)
	db.Model(&mDokter).Update("nik", dokter.Nik)
	db.Model(&mDokter).Update("alamat", dokter.Alamat)
	c.JSON(200, gin.H{
		"code": 200,
		"data": dokter,
	})
}
