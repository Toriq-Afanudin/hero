package controller

import (
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
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
	Jenis_penyakit   string `json:"jenis_penyakit"`
	Jenis_penanganan string `json:"jenis_penanganan"`
}

func Tambah_data_pasien(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	db := c.MustGet("db").(*gorm.DB)
	var user model.User
	db.Where("email = ?", claims["id"]).Where("level = ?", "admin").Find(&user)
	if claims["id"] == user.Email {
		c.JSON(400, gin.H{
			"status":  "Error",
			"message": "Halaman ini hanya bisa diakses oleh dokter atau perawat.",
		})
		return
	}
	var t tambah_data
	if err := c.ShouldBindJSON(&t); err != nil {
		c.JSON(400, gin.H{
			"status":  "Error",
			"message": "Request harus dalam bentuk JSON.",
		})
		return
	}
	add := model.Pasien{
		Nik:           t.Nik,
		Nama:          t.Nama,
		Alamat:        t.Alamat,
		Jenis_kelamin: t.Jenis_kelamin,
		No_hp:         t.Nomer_telfon,
		Tempat_lahir:  t.Tempat_lahir,
		Tanggal_lahir: t.Tanggal_lahir,
	}

	if (t.Nik == "") || (t.Nama == "") || (t.Alamat == "") || (t.Jenis_kelamin == "") || (t.Nomer_telfon == "") || (t.Tempat_lahir == "") || (t.Tanggal_lahir == "") {
		c.JSON(400, gin.H{
			"status":  "Error",
			"message": "Tidak boleh ada data yang kosong.",
		})
		return
	}
	var pasien model.Pasien
	db.Where("nik = ?", t.Nik).Find(&pasien)
	if t.Nik == pasien.Nik {
		c.JSON(200, gin.H{
			"status":  "Berhasil",
			"message": "NIK sudah tercantum dalam sistem, akan ditambahkan data rekam medis.",
		})
		var psn model.Pasien
		db.Where("nama = ?", t.Nama).Find(&psn)
		add2 := model.Rekam_medis{
			Tanggal:          time.Now(),
			Pemeriksaan:      t.Jenis_penyakit,
			Jenis_penanganan: t.Jenis_penanganan,
			Id_pasien:        psn.Id,
		}
		db.Create(&add2)
		if t.Jenis_penanganan == "rawat jalan" {
			tambah_rawat_jalan := model.Rawat_jalan{
				Id: psn.Id,
			}
			db.Create(&tambah_rawat_jalan)
		}
		return
	}
	if (t.Jenis_kelamin == "P") || (t.Jenis_kelamin == "L") {
		db.Create(&add)
		c.JSON(200, gin.H{
			"status":  "Berhasil",
			"data":    add,
			"user":    claims["id"],
			"message": "Lengkapi data rekam medis.",
		})
		var psn model.Pasien
		db.Where("nama = ?", t.Nama).Find(&psn)
		add2 := model.Rekam_medis{
			Tanggal:          time.Now(),
			Pemeriksaan:      t.Jenis_penyakit,
			Jenis_penanganan: t.Jenis_penanganan,
			Id_pasien:        psn.Id,
		}
		db.Create(&add2)
		if t.Jenis_penanganan == "rawat jalan" {
			tambah_rawat_jalan := model.Rawat_jalan{
				Id: psn.Id,
			}
			db.Create(&tambah_rawat_jalan)
		}
	} else {
		c.JSON(400, gin.H{
			"status":  "Error",
			"message": "Jenis kelamin harus di isi dengan L atau P",
		})
	}

}
