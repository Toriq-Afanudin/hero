package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"heroku.com/model"
)

type edit_data struct {
	Nik           string `json:"nik"`
	Nama          string `json:"nama"`
	Alamat        string `json:"alamat"`
	Jenis_kelamin string `json:"jenis_kelamin"`
	Nomer_telfon  string `json:"nomer_telfon"`
	Tempat_lahir  string `json:"tempat_lahir"`
	Tanggal_lahir string `json:"tanggal_lahir"`
}

func Edit_data_pasien(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var edit edit_data
	if err := c.ShouldBindJSON(&edit); err != nil {
		c.JSON(400, gin.H{
			"status":  "Error",
			"message": "Request harus dalam bentuk JSON.",
		})
		return
	}
	var pasien model.Pasien
	db.Where("id = ?", c.Param("id")).Find(&pasien)
	if pasien.Nama == "" {
		c.JSON(400, gin.H{
			"status":  "Error",
			"message": "Parameter id tidak ditemukan.",
		})
		return
	}
	db.Model(&pasien).Update("nik", edit.Nik)
	db.Model(&pasien).Update("nama", edit.Nama)
	db.Model(&pasien).Update("alamat", edit.Alamat)
	db.Model(&pasien).Update("jenis_kelamin", edit.Jenis_kelamin)
	db.Model(&pasien).Update("no_hp", edit.Nomer_telfon)
	db.Model(&pasien).Update("tempat_lahir", edit.Tempat_lahir)
	db.Model(&pasien).Update("tanggal_lahir", edit.Tanggal_lahir)
	c.JSON(200, gin.H{
		"code": 200,
		"data": pasien,
	})
}
