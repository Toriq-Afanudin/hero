package controller

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"heroku.com/model"
)

type tambah_data struct {
	Nik           string `json:"nik"`
	Nama          string `json:"nama"`
	Jenis_kelamin string `json:"jenis_kelamin"`
	Poli          string `json:"poli"`
	Alamat        string `json:"alamat"`
	No_hp         string `json:"no_hp"`
	Tempat_lahir  string `json:"tempat_lahir"`
	Tanggal_lahir string `json:"tanggal_lahir"`
}

func Tambah_data_pasien(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var data tambah_data
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"message": "Request harus dalam bentuk JSON.",
		})
		return
	}
	if (data.Nik == "") || (data.Nama == "") || (data.Alamat == "") || (data.Jenis_kelamin == "") || (data.No_hp == "") || (data.Tempat_lahir == "") || (data.Tanggal_lahir == "") || (data.Poli == "") {
		c.JSON(400, gin.H{
			"code":    400,
			"message": "Tidak boleh ada data yang kosong.",
		})
		return
	}
	var pasien model.Pasien
	db.Where("nik = ?", data.Nik).Find(&pasien)
	if data.Nik == pasien.Nik {
		c.JSON(400, gin.H{
			"code":    400,
			"message": "NIK sudah digunakan, pasien telah terdaftar.",
		})
		return
	}
	var poli []string
	poli = append(poli, "Gigi", "Kandungan", "THT", "Umum")
	if (data.Poli == poli[0]) || (data.Poli == poli[1]) || (data.Poli == poli[2]) || (data.Poli == poli[3]) {
		fmt.Println("poli benar")
	} else {
		c.JSON(400, gin.H{
			"code":    400,
			"message": "Poli yang tersedia: Gigi, Kandungan, THT, dan Umum.",
		})
		return
	}
	tambah := model.Pasien{
		Nik:           data.Nik,
		Nama:          data.Nama,
		Alamat:        data.Alamat,
		Jenis_kelamin: data.Jenis_kelamin,
		No_hp:         data.No_hp,
		Tempat_lahir:  data.Tempat_lahir,
		Tanggal_lahir: data.Tanggal_lahir,
	}
	db.Create(&tambah)
	db.Where("nik = ?", data.Nik).Find(&pasien)
	var rekam model.Rekam_medis
	db.Where("poli = ?", data.Poli).Find(&rekam)
	nAntri := rekam.Nomor_antrian + 1
	rekam_medis := model.Rekam_medis{
		Id_pasien:     pasien.Id,
		Tanggal:       time.Now(),
		Poli:          data.Poli,
		Nomor_antrian: nAntri,
	}
	db.Create(&rekam_medis)
	c.JSON(200, gin.H{
		"code":    200,
		"data":    tambah,
		"message": "Lengkapi data rekam medis.",
	})
}
