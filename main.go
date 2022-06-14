package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"heroku.com/controller"
)

func main() {
	r := gin.Default()
	r.GET("/", controller.Utama)

	godotenv.Load()
	port := os.Getenv("PORT")
	address := fmt.Sprintf("%s:%s", "0.0.0.0", port)
	fmt.Println(address)

	r.Run(address)
}
