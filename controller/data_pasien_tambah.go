package controller

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"heroku.com/model"
)

type tambah_data struct {
	Nik           string `json:"nik"`
	Nama          string `json:"nama"`
	Alamat        string `json:"alamat"`
	Jenis_kelamin string `json:"jenis_kelamin"`
	Nomer_telfon  string `json:"nomer_telfon"`
	Tempat_lahir  string `json:"tempat_lahir"`
	Tanggal_lahir string `json:"tanggal_lahir"`
}

func Tambah_data_pasien(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	db := c.MustGet("db").(*gorm.DB)
	var user model.User
	db.Where("email = ?", claims["id"]).Where("level = ?", "admin").Find(&user)
	if claims["id"] == user.Email {
		c.JSON(400, gin.H{
			"status":  "gagal menampilkan data",
			"message": "yang berhak mengakses halaman ini hanya dokter atau perawat",
		})
		return
	}
	var t tambah_data
	if err := c.ShouldBindJSON(&t); err != nil {
		c.JSON(400, gin.H{
			"status":  "error",
			"message": "input tidak dalam bentuk json",
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
			"status": "gagal menambahkan, tidak boleh ada data yang kosong",
		})
	}
	if (t.Jenis_kelamin == "P") || (t.Jenis_kelamin == "L") {
		db.Create(&add)
		c.JSON(200, gin.H{
			"status": "berhasil menambahkan data pasien",
			"data":   add,
			"userID": claims["id"],
		})
	} else {
		c.JSON(400, gin.H{
			"status": "gagal menambahkan, jenis kelamin harus di isi dengan P atau L",
		})
	}

}
