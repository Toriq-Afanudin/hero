package controller

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"heroku.com/model"
)

type rawat_jalan struct {
	Id                 int    `json:"id"`
	Nik                string `json:"nik"`
	Nama               string `json:"nama"`
	Alamat             string `json:"alamat"`
	Jenis_kelamin      string `json:"jenis_kelamin"`
	No_hp              string `json:"no_hp"`
	Tempat_lahir       string `json:"tempat_lahir"`
	Tanggal_lahir      string `json:"tanggal_lahir"`
	Jenis_penyakit     string `json:"jenis_penyakit"`
	Jenis_penanganan   string `json:"jenis_penanganan"`
	Jadwal_rawat_jalan string `json:"jadwal_rawat_jalan"`
	Nomer_antrian      string `json:"nomer_antrian"`
}

func Rawat_jalan_lihat(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	db := c.MustGet("db").(*gorm.DB)
	var user model.User
	db.Where("email = ?", claims["id"]).Where("level = ?", "admin").Find(&user)
	if claims["id"] == user.Email {
		c.JSON(400, gin.H{
			"status":  "Error",
			"message": "Halaman ini hanya bisa diakses oleh dokter atau perawat.",
		})
		return
	}
	var rkm_medis []model.Rekam_medis
	var rawat []interface{}
	db.Where("jenis_penanganan = ?", "rawat jalan").Find(&rkm_medis)
	for i := 0; i < len(rkm_medis); i++ {
		var psn model.Pasien
		var r_jalan model.Rawat_jalan
		db.Where("id = ?", rkm_medis[i].Id_pasien).Find(&psn)
		db.Where("id = ?", rkm_medis[i].Id_pasien).Find(&r_jalan)
		new := rawat_jalan{
			Id:                 psn.Id,
			Nik:                psn.Nik,
			Nama:               psn.Nama,
			Alamat:             psn.Alamat,
			Jenis_kelamin:      psn.Jenis_kelamin,
			No_hp:              psn.No_hp,
			Tempat_lahir:       psn.Tempat_lahir,
			Tanggal_lahir:      psn.Tanggal_lahir,
			Jenis_penyakit:     rkm_medis[i].Pemeriksaan,
			Jenis_penanganan:   rkm_medis[i].Jenis_penanganan,
			Jadwal_rawat_jalan: r_jalan.Jadwal_rawat_jalan,
			Nomer_antrian:      r_jalan.Nomer_antrian,
		}
		rawat = append(rawat, new)
	}
	c.JSON(200, gin.H{
		"status": "Berhasil",
		"data":   rawat,
		"user":   claims["id"],
	})
}
