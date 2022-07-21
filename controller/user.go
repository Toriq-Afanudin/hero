package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"heroku.com/model"
)

type GetUser struct {
	Id            int    `json:"id"`
	Sip           string `json:"sip"`
	Nama          string `json:"nama"`
	Jenis_kelamin string `json:"jenis_kelamin"`
	Email         string `json:"email"`
	Password      string `json:"password"`
	Level         string `json:"level"`
}

type InputUser struct {
	Nama     string `json:"nama"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Level    string `json:"level"`
}

func GetDataUser(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var user []model.User
	db.Find(&user)
	var GetData []interface{}
	for i := 0; i < len(user); i++ {
		if user[i].Level == "Perawat" {
			var perawat model.Perawat
			db.Where("id_user = ?", user[i].Id).Find(&perawat)
			getUser := GetUser{
				Id:            user[i].Id,
				Sip:           perawat.Sip,
				Nama:          user[i].Nama,
				Jenis_kelamin: perawat.Jenis_kelamin,
				Email:         user[i].Email,
				Password:      user[i].Password,
				Level:         user[i].Level,
			}
			GetData = append(GetData, getUser)
		}
		if user[i].Level == "Dokter" {
			var dokter model.Dokter
			db.Where("id_user = ?", user[i].Id).Find(&dokter)
			getUser := GetUser{
				Id:            user[i].Id,
				Sip:           dokter.Sip,
				Nama:          user[i].Nama,
				Jenis_kelamin: dokter.Jenis_kelamin,
				Email:         user[i].Email,
				Password:      user[i].Password,
				Level:         user[i].Level,
			}
			GetData = append(GetData, getUser)
		}
		if user[i].Level == "Admin" {
			continue
		}
	}
	c.JSON(200, gin.H{
		"code": 200,
		"data": GetData,
	})
}

// func GetDataUser(c *gin.Context) {
// 	db := c.MustGet("db").(*gorm.DB)
// 	var user []GetUser
// 	db.Raw("select * from users order by id desc").Scan(&user)
// 	for i := 0; i < len(user); i++ {
// 		if user[i].Level == "perawat" {
// 			var perawat model.Perawat
// 			db.Raw("select sip, nama_perawat, jenis_kelamin from perawats where id_user=?;", user[i].Id).Scan(&perawat)
// 			db.Model(&user[i]).Update("sip", perawat.Sip)
// 			// db.Model(&user[i]).Update("nama", perawat.Nama_perawat)
// 			db.Model(&user[i]).Update("jenis_kelamin", perawat.Jenis_kelamin)
// 		}
// 		if user[i].Level == "dokter" {
// 			var dokter model.Dokter
// 			db.Raw("select sip, nama_dokter, jenis_kelamin from dokters where id_user=?;", user[i].Id).Scan(&dokter)
// 			db.Model(&user[i]).Update("sip", dokter.Sip)
// 			// db.Model(&user[i]).Update("nama", dokter.Nama_dokter)
// 			db.Model(&user[i]).Update("jenis_kelamin", dokter.Jenis_kelamin)
// 		}
// 	}
// 	c.JSON(200, gin.H{
// 		"code": 200,
// 		"data": user,
// 	})
// }

func PostDataUser(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var input InputUser
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"message": "Request harus dalam bentuk JSON.",
		})
		return
	}
	if (input.Nama == "") || (input.Email == "") || (input.Password == "") || (input.Level == "") {
		c.JSON(400, gin.H{
			"code":    400,
			"message": "Data tidak boleh ada yang kosong.",
		})
		return
	}
	var user model.User
	db.Where("nickname = ?", input.Nama).Find(&user)
	if user.Nama == input.Nama {
		c.JSON(400, gin.H{
			"code":    400,
			"message": "Nickname sudah digunakan.",
		})
		return
	}
	db.Where("email = ?", input.Email).Find(&user)
	if user.Email == input.Email {
		c.JSON(400, gin.H{
			"code":    400,
			"message": "Email sudah digunakan.",
		})
		return
	}
	if len(input.Password) < 8 {
		c.JSON(400, gin.H{
			"code":    400,
			"message": "Password minimal terdiri dari 8 karakter.",
		})
		return
	}
	var pekerjaan bool
	if (input.Level == "Perawat") || (input.Level == "Dokter") || (input.Level == "Admin") {
		pekerjaan = true
	} else {
		c.JSON(400, gin.H{
			"code":    400,
			"message": "Pilih jenis pekerjaan sebagai Admin, Dokter, atau Perawat. Penulisan diawal huruf kapital",
		})
		return
	}
	create := model.User{
		Nama:     input.Nama,
		Email:    input.Email,
		Password: input.Password,
		Level:    input.Level,
	}
	if pekerjaan {
		db.Create(&create)
		var mod model.User
		db.Where("nama = ?", input.Nama).Find(&mod)
		var message string
		if input.Level == "Dokter" {
			new := model.Dokter{
				Id_user: mod.Id,
			}
			db.Create(&new)
			message = "Dokter " + input.Nama + " diharuskan melengkapi data sendiri"
		}
		if input.Level == "Perawat" {
			new := model.Perawat{
				Id_user: mod.Id,
			}
			db.Create(&new)
			message = "Perawat " + input.Nama + " diharuskan melengkapi data sendiri"
		}
		if input.Level == "Admin" {
			message = input.Nama + " resmi menjadi admin baru"
		}
		c.JSON(200, gin.H{
			"code":    200,
			"data":    input,
			"message": message,
		})
	}
}

func DeleteDataUser(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var user model.User
	db.Where("id = ?", c.Param("id")).Find(&user)
	if user.Id == 0 {
		c.JSON(400, gin.H{
			"code":    400,
			"data":    "-",
			"message": "Parameter id yang anda masukan salah.",
		})
		return
	}
	db.Where("id = ?", c.Param("id")).Delete(&user)
	c.JSON(200, gin.H{
		"code":    200,
		"data":    user,
		"message": "Data user berhasil dihapus.",
	})
}
