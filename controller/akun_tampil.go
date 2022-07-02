package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"heroku.com/model"
)

type user struct {
	Id            string `json:"id"`
	Sip           string `json:"sip"`
	Nama          string `json:"nama"`
	Jenis_kelamin string `json:"jenis_kelamin"`
	Email         string `json:"email"`
	Password      string `json:"password"`
	Level         string `json:"level"`
}

func Akun_tampil(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var akun []user
	db.Raw("SELECT id, email, password, level FROM users WHERE level=? OR level=?", "perawat", "dokter").Scan(&akun)
	for i := 0; i < len(akun); i++ {
		if akun[i].Level == "perawat" {
			var perawat model.Perawat
			db.Raw("SELECT sip, nama_perawat, jenis_kelamin FROM perawats WHERE id_user=?;", akun[i].Id).Scan(&perawat)
			db.Model(&akun[i]).Update("sip", perawat.Sip)
			db.Model(&akun[i]).Update("nama", perawat.Nama_perawat)
			db.Model(&akun[i]).Update("jenis_kelamin", perawat.Jenis_kelamin)
		}
		if akun[i].Level == "dokter" {
			var dokter model.Dokter
			db.Raw("SELECT sip, nama_dokter, jenis_kelamin FROM dokters WHERE id_user=?;", akun[i].Id).Scan(&dokter)
			db.Model(&akun[i]).Update("sip", dokter.Sip)
			db.Model(&akun[i]).Update("nama", dokter.Nama_dokter)
			db.Model(&akun[i]).Update("jenis_kelamin", dokter.Jenis_kelamin)
		}
	}
	c.JSON(200, gin.H{
		"code": 200,
		"data": akun,
	})
}
