package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"heroku.com/model"
)

type per struct {
	Id            int
	Nama          string
	Sip           string
	Jenis_kelamin string
	Poli          string
	Jadwal_kerja  string
	Jabatan       string
	Nomor_telfon  string
	Nomor_str     string
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
