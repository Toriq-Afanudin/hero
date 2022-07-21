package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"heroku.com/model"
)

func Data_perawat_lihat(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var perawat []model.Perawat
	db.Raw("select * from perawats order by id_user desc").Scan(&perawat)
	c.JSON(200, gin.H{
		"code": 200,
		"data": perawat,
	})
}
