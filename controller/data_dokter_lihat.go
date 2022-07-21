package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"heroku.com/model"
)

type dok struct {
	Id             int
	Nama           string
	Sip            string
	Jenis_kelamin  string
	Jadwal_praktek string
	Poli           string
	Nomor_telfon   string
	Nomor_str      string
}

func Data_dokter_lihat(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var dokter []model.Dokter
	db.Raw("select * from dokters order by id_user desc").Scan(&dokter)
	var GetDokter []interface{}
	for i := 0; i < len(dokter); i++ {
		var user model.User
		db.Where("id = ?", dokter[i].Id_user).Find(&user)
		new := dok{
			Id:             dokter[i].Id_user,
			Nama:           user.Nama,
			Sip:            dokter[i].Sip,
			Jenis_kelamin:  dokter[i].Jenis_kelamin,
			Jadwal_praktek: dokter[i].Jadwal_praktek,
			Poli:           dokter[i].Poli,
			Nomor_telfon:   dokter[i].Nomor_telfon,
			Nomor_str:      dokter[i].Nomor_str,
		}
		GetDokter = append(GetDokter, new)
	}
	c.JSON(200, gin.H{
		"code": 200,
		"data": GetDokter,
	})
}
