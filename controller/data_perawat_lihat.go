package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"heroku.com/model"
)

type per struct {
	Id            int    `json:"id"`
	Nama          string `json:"nama"`
	Sip           string `json:"sip"`
	Jenis_kelamin string `json:"jenis_kelamin"`
	Poli          string `json:"poli"`
	Jadwal_kerja  string `json:"jadwal_kerja"`
	Jabatan       string `json:"jabatan"`
	Nomor_telfon  string `json:"nomor_telfon"`
	Nomor_str     string `json:"nomor_str"`
}

func Data_perawat_lihat(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var perawat []model.Perawat
	db.Raw("select * from perawats order by id_user desc").Scan(&perawat)
	var GetPerawat []interface{}
	for i := 0; i < len(perawat); i++ {
		var user model.User
		db.Where("id = ?", perawat[i].Id_user).Find(&user)
		new := per{
			Id:            perawat[i].Id_user,
			Nama:          user.Nama,
			Sip:           perawat[i].Sip,
			Jenis_kelamin: perawat[i].Jenis_kelamin,
			Poli:          perawat[i].Poli,
			Jadwal_kerja:  perawat[i].Jadwal_kerja,
			Jabatan:       perawat[i].Jabatan,
			Nomor_telfon:  perawat[i].Nomor_telfon,
			Nomor_str:     perawat[i].Nomor_str,
		}
		GetPerawat = append(GetPerawat, new)
	}
	c.JSON(200, gin.H{
		"code": 200,
		"data": GetPerawat,
	})
}
