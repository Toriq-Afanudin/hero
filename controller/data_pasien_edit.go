package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"heroku.com/model"
)

type edit_data struct {
	Nik           string `json:"nik"`
	Nama          string `json:"nama"`
	Jenis_kelamin string `json:"jenis_kelamin"`
	Poli          string `json:"poli"`
	Alamat        string `json:"alamat"`
	No_hp         string `json:"no_hp"`
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
	if (edit.Poli == "Gigi") || (edit.Poli == "Kandungan") || (edit.Poli == "THT") || (edit.Poli == "Umum") {
	} else {
		if pasien.Nama == "" {
			c.JSON(400, gin.H{
				"status":  "Error",
				"message": "Poli yang tersedia: Gigi, THT, Kandungan, dan Umum",
			})
			return
		}
	}
	var rekam model.Rekam_medis
	db.Model(&pasien).Where("id = ?", c.Param("id")).Update("nik", edit.Nik)
	db.Model(&pasien).Where("id = ?", c.Param("id")).Update("nama", edit.Nama)
	db.Model(&pasien).Where("id = ?", c.Param("id")).Update("jenis_kelamin", edit.Jenis_kelamin)
	db.Model(&rekam).Where("id_pasien = ?", c.Param("id")).Update("poli", edit.Poli)
	db.Model(&pasien).Where("id = ?", c.Param("id")).Update("alamat", edit.Alamat)
	db.Model(&pasien).Where("id = ?", c.Param("id")).Update("no_hp", edit.No_hp)
	db.Model(&pasien).Where("id = ?", c.Param("id")).Update("tempat_lahir", edit.Tempat_lahir)
	db.Model(&pasien).Where("id = ?", c.Param("id")).Update("tanggal_lahir", edit.Tanggal_lahir)
	c.JSON(200, gin.H{
		"code": 200,
		"data": pasien,
	})
}
