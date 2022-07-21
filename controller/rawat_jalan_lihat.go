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
	Proses             bool   `json:"proses"`
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
	db.Raw("select * from rawat_jalans order by id_pasien desc").Scan(&rJalan)
	var daftarRawatJalan []interface{}
	for i := 0; i < len(rJalan); i++ {
		var pasien model.Pasien
		db.Where("id = ?", rJalan[i].Id_pasien).Find(&pasien)
		var rMedis model.Rekam_medis
		db.Where("id_pasien = ?", rJalan[i].Id_pasien).Find(&rMedis)
		var poli string
		if rJalan[i].Poli == "Gigi" {
			poli = "G-"
		}
		if rJalan[i].Poli == "Kandungan" {
			poli = "K-"
		}
		if rJalan[i].Poli == "THT" {
			poli = "T-"
		}
		if rJalan[i].Poli == "Umum" {
			poli = "U-"
		}
		var nomor = rJalan[i].Nomer_antrian
		var str = strconv.Itoa(nomor)
		nAntrian := poli + str
		new := rawat_jalan{
			Id:                 pasien.Id,
			Nik:                pasien.Nik,
			Nama:               pasien.Nama,
			Alamat:             pasien.Alamat,
			Jenis_kelamin:      pasien.Jenis_kelamin,
			Jadwal_rawat_jalan: rJalan[i].Jadwal_rawat_jalan,
			Nomer_antrian:      nAntrian,
			Proses:             rJalan[i].Bool,
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

func Rawat_jalan_lihat_per_poli(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var dokter model.Dokter
	var poli string
	db.Where("id_user = ?", c.Param("id")).Find(&dokter)
	poli = dokter.Poli
	if dokter.Poli == "" {
		var perawat model.Perawat
		db.Where("id_user = ?", c.Param("id")).Find(&perawat)
		poli = perawat.Poli
		if perawat.Poli == "" {
			c.JSON(400, gin.H{
				"code":    400,
				"message": "Parameter id yang anda masukan salah",
			})
			return
		}
	}
	var rJalan []model.Rawat_jalan
	db.Raw("select * from rawat_jalans where poli=? order by id_pasien desc", poli).Scan(&rJalan)
	var daftarRawatJalan []interface{}
	for i := 0; i < len(rJalan); i++ {
		var pasien model.Pasien
		db.Where("id = ?", rJalan[i].Id_pasien).Find(&pasien)
		var rMedis model.Rekam_medis
		db.Where("id_pasien = ?", rJalan[i].Id_pasien).Find(&rMedis)
		var poli string
		if rJalan[i].Poli == "Gigi" {
			poli = "G-"
		}
		if rJalan[i].Poli == "Kandungan" {
			poli = "K-"
		}
		if rJalan[i].Poli == "THT" {
			poli = "T-"
		}
		if rJalan[i].Poli == "Umum" {
			poli = "U-"
		}
		var nomor = rJalan[i].Nomer_antrian
		var str = strconv.Itoa(nomor)
		nAntrian := poli + str
		new := rawat_jalan{
			Id:                 pasien.Id,
			Nik:                pasien.Nik,
			Nama:               pasien.Nama,
			Alamat:             pasien.Alamat,
			Jenis_kelamin:      pasien.Jenis_kelamin,
			Jadwal_rawat_jalan: rJalan[i].Jadwal_rawat_jalan,
			Nomer_antrian:      nAntrian,
			Proses:             rJalan[i].Bool,
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
