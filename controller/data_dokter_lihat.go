package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"heroku.com/model"
)

func Data_dokter_lihat(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var dokter []model.Dokter
	db.Raw("select * from dokters order by id_user desc").Scan(&dokter)
	c.JSON(200, gin.H{
		"code": 200,
		"data": dokter,
	})
}
