package controller

import (
	"strconv"

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
	Jadwal_rawat_jalan string `json:"jadwal_rawat_jalan"`
	Nomer_antrian      string `json:"nomer_antrian"`
	Proses             string `json:"proses"`
	Keterangan         string `json:"keterangan"`
	No_hp              string `json:"no_hp"`
	Tempat_lahir       string `json:"tempat_lahir"`
	Tanggal_lahir      string `json:"tanggal_lahir"`
	Jenis_penyakit     string `json:"jenis_penyakit"`
	Jenis_penanganan   string `json:"jenis_penanganan"`
}

func Rawat_jalan_lihat(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var rJalan []model.Rawat_jalan
	db.Where("poli = ?", c.Param("poli")).Find(&rJalan)
	var daftarRawatJalan []interface{}
	for i := 0; i < len(rJalan); i++ {
		var pasien model.Pasien
		db.Where("id = ?", rJalan[i].Id).Find(&pasien)
		var rMedis model.Rekam_medis
		db.Where("id_pasien = ?", rJalan[i].Id).Find(&rMedis)
		var poli string
		if rJalan[i].Poli == "gigi" {
			poli = "G-"
		}
		if rJalan[i].Poli == "kulit" {
			poli = "K-"
		}
		if rJalan[i].Poli == "tht" {
			poli = "T-"
		}
		if rJalan[i].Poli == "umum" {
			poli = "U-"
		}
		var nomor = rJalan[i].Nomer_antrian
		var str = strconv.Itoa(nomor)
		nAntrian := poli + str
		var proses string
		if rJalan[i].Bool == 0 {
			proses = "antre"
		} else {
			proses = "done"
		}
		new := rawat_jalan{
			Id:                 pasien.Id,
			Nik:                pasien.Nik,
			Nama:               pasien.Nama,
			Alamat:             pasien.Alamat,
			Jenis_kelamin:      pasien.Jenis_kelamin,
			Jadwal_rawat_jalan: rJalan[i].Jadwal_rawat_jalan,
			Nomer_antrian:      nAntrian,
			Proses:             proses,
			Keterangan:         rJalan[i].Keterangan,
			No_hp:              pasien.No_hp,
			Tempat_lahir:       pasien.Tempat_lahir,
			Tanggal_lahir:      pasien.Tanggal_lahir,
			Jenis_penyakit:     rMedis.Pemeriksaan,
			Jenis_penanganan:   rMedis.Jenis_penanganan,
		}
		daftarRawatJalan = append(daftarRawatJalan, new)
	}
	c.JSON(200, gin.H{
		"code": 200,
		"data": daftarRawatJalan,
	})
}
