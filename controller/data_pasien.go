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
	Jenis_kelamin string `json:"jenis_kelamin"`
	Poli          string `json:"poli"`
	Nama_dokter   string `json:"nama_dokter"`
	Alamat        string `json:"alamat"`
	No_hp         string `json:"no_hp"`
	Tempat_lahir  string `json:"tempat_lahir"`
	Tanggal_lahir string `json:"tanggal_lahir"`
}

func DataPasien(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var mPasien []model.Pasien
	db.Find(&mPasien)
	var daftar_pasien []interface{}
	for i := 0; i < len(mPasien); i++ {
		var rekam_medis model.Rekam_medis
		db.Where("id_pasien = ?", mPasien[i].Id).Find(&rekam_medis)
		var dokter model.Dokter
		db.Where("poli = ?", rekam_medis.Poli).Find(&dokter)
		new := pasien{
			Id:            mPasien[i].Id,
			Nik:           mPasien[i].Nik,
			Nama:          mPasien[i].Nama,
			Jenis_kelamin: mPasien[i].Jenis_kelamin,
			Poli:          rekam_medis.Poli,
			Nama_dokter:   dokter.Nama_dokter,
			Alamat:        mPasien[i].Alamat,
			No_hp:         mPasien[i].No_hp,
			Tempat_lahir:  mPasien[i].Tempat_lahir,
			Tanggal_lahir: mPasien[i].Tanggal_lahir,
		}
		daftar_pasien = append(daftar_pasien, new)
	}
	c.JSON(200, gin.H{
		"code": 200,
		"data": daftar_pasien,
	})
}

// var antri string
// if rekam_medis.Poli == "Gigi" {
// 	antri = "G-"
// }
// if rekam_medis.Poli == "Kandungan" {
// 	antri = "K-"
// }
// if rekam_medis.Poli == "THT" {
// 	antri = "T-"
// }
// if rekam_medis.Poli == "Umum" {
// 	antri = "U-"
// }
// var str = strconv.Itoa(rekam_medis.Nomor_antrian)
// nAntrian := antri + str
