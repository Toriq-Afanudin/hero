package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"heroku.com/model"
)

type pasien struct {
	Id            int    `json:"id"`
	Nik           string `json:"nik"`
	Nama          string `json:"nama"`
	Alamat        string `json:"alamat"`
	Jenis_kelamin string `json:"jenis_kelamin"`
	No_hp         string `json:"no_hp"`
	Tempat_lahir  string `json:"tempat_lahir"`
	Tanggal_lahir string `json:"tanggal_lahir"`
	Rekam_medis   interface{}
}

type rekam_medis struct {
	Tanggal          string `json:"tanggal"`
	Keluhan          string `json:"keluhan"`
	Poli             string `json:"poli"`
	Pemeriksaan      string `json:"pemeriksaan"`
	Jenis_penanganan string `json:"jenis_penanganan"`
}

func DataPasien(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var mPasien []model.Pasien
	db.Find(&mPasien)
	var data []interface{}
	for i := 0; i < len(mPasien); i++ {
		var rekam_medis []rekam_medis
		db.Raw("select tanggal, keluhan, poli, pemeriksaan, jenis_penanganan from rekam_medis where id_pasien=?", mPasien[i].Id).Scan(&rekam_medis)
		tampil := pasien{
			Id:            mPasien[i].Id,
			Nik:           mPasien[i].Nik,
			Nama:          mPasien[i].Nama,
			Alamat:        mPasien[i].Alamat,
			Jenis_kelamin: mPasien[i].Jenis_kelamin,
			No_hp:         mPasien[i].No_hp,
			Tempat_lahir:  mPasien[i].Tempat_lahir,
			Tanggal_lahir: mPasien[i].Tanggal_lahir,
			Rekam_medis:   rekam_medis,
		}
		data = append(data, tampil)
	}
	c.JSON(200, gin.H{
		"code": 200,
		"data": data,
	})
}
