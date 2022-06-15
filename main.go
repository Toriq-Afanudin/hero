package main

import (
	"fmt"
	"os"

	// "database/sql"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"heroku.com/controller"
	"heroku.com/model"
)

func main() {
	r := gin.Default()
	db := model.SetupModels()
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})
	r.GET("/", controller.Utama)
	r.GET("/user", controller.GetUser)
	r.GET("/login", controller.Login)

	godotenv.Load()
	ip := os.Getenv("IP")
	port := os.Getenv("PORT2")
	address := fmt.Sprintf("%s:%s", ip, port)
	fmt.Println(address)

	r.Run(address)

	// r.Run()
}
