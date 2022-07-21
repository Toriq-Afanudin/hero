package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"heroku.com/model"
)

type dok struct {
	Id             int    `json:"id"`
	Nama           string `json:"nama"`
	Sip            string `json:"sip"`
	Jenis_kelamin  string `json:"jenis_kelamin"`
	Jadwal_praktek string `json:"jadwal_praktek"`
	Poli           string `json:"poli"`
	Nomor_telfon   string `json:"nomor_telfon"`
	Nomor_str      string `json:"nomor_str"`
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
