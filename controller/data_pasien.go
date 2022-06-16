package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type pasien struct {
	Id             int    `form:"id"`
	Nik            string `form:"nik"`
	Nama           string `form:"nama"`
	Alamat         string `form:"alamat"`
	Jenis_kelamin  string `form:"jenis_kelamin"`
	Jenis_penyakit interface{}
}

type penyakit struct {
	Pemeriksaan string `form:"pemeriksaan"`
}

func DataPasien(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var d []pasien
	db.Raw("SELECT id, nik, nama, alamat, jenis_kelamin FROM capstone.pasiens;").Scan(&d)
	var p penyakit
	for i := 0; i < len(d); i++ {
		db.Raw("SELECT pemeriksaan FROM capstone.rekam_medis WHERE id=?", d[i].Id).Scan(&p)
		db.Model(&d[i]).Update("Jenis_penyakit", p.Pemeriksaan)
	}
	c.JSON(200, gin.H{
		"status": "berhasil menampilkan data pasien",
		"akun":   d,
	})
}
