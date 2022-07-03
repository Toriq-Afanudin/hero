package controller

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"heroku.com/model"
)

type tambah_data struct {
	Nik              string `json:"nik"`
	Nama             string `json:"nama"`
	Alamat           string `json:"alamat"`
	Jenis_kelamin    string `json:"jenis_kelamin"`
	Nomer_telfon     string `json:"nomer_telfon"`
	Tempat_lahir     string `json:"tempat_lahir"`
	Tanggal_lahir    string `json:"tanggal_lahir"`
	Poli             string `json:"poli"`
	Jenis_penanganan string `json:"jenis_penanganan"`
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
	if (data.Nik == "") || (data.Nama == "") || (data.Alamat == "") || (data.Jenis_kelamin == "") || (data.Nomer_telfon == "") || (data.Tempat_lahir == "") || (data.Tanggal_lahir == "") || (data.Poli == "") || (data.Jenis_penanganan == "") {
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
	poli = append(poli, "gigi", "kulit", "tht", "umum")
	if (data.Poli == poli[0]) || (data.Poli == poli[1]) || (data.Poli == poli[2]) || (data.Poli == poli[3]) {
		fmt.Println("poli benar")
	} else {
		c.JSON(400, gin.H{
			"code":    400,
			"message": "Poli yang tersedia: gigi, kulit, tht, dan umum.",
		})
		return
	}
	tambah := model.Pasien{
		Nik:           data.Nik,
		Nama:          data.Nama,
		Alamat:        data.Alamat,
		Jenis_kelamin: data.Jenis_kelamin,
		No_hp:         data.Nomer_telfon,
		Tempat_lahir:  data.Tempat_lahir,
		Tanggal_lahir: data.Tanggal_lahir,
	}
	db.Create(&tambah)
	db.Where("nik = ?", data.Nik).Find(&pasien)
	rekam_medis := model.Rekam_medis{
		Id_pasien:        pasien.Id,
		Tanggal:          time.Now(),
		Poli:             data.Poli,
		Jenis_penanganan: data.Jenis_penanganan,
	}
	db.Create(&rekam_medis)
	if data.Jenis_penanganan == "rawat jalan" {
		var no int
		for i := 0; i < len(poli); i++ {
			if data.Poli == poli[i] {
				var mRJalan model.Rawat_jalan
				db.Where("poli = ?", poli[i]).Where("bool = ?", 0).Find(&mRJalan)
				no = mRJalan.Nomer_antrian + 1
				break
			}
		}
		rawat_jalan := model.Rawat_jalan{
			Id_pasien:     pasien.Id,
			Tanggal:       time.Now(),
			Poli:          data.Poli,
			Nomer_antrian: no,
		}
		db.Create(&rawat_jalan)
		c.JSON(200, gin.H{
			"code":    200,
			"data":    tambah,
			"message": "Lengkapi data rekam medis dan rawat jalan.",
		})
		return
	}
	c.JSON(200, gin.H{
		"code":    200,
		"data":    tambah,
		"message": "Lengkapi data rekam medis.",
	})
}
