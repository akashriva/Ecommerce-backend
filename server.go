package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"

	"github.com/akashriva/gin_framework/config"
	"github.com/akashriva/gin_framework/middlewares"
)

func init() {
	

	config.InitDdConnection()
}

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	router := gin.New()
	//middlewares
	router.Use(gin.Recovery(), middlewares.Logger())

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "server contect",
		})
	})

	log.Printf("Server is connected on Port : %s", port)

	router.Run(":" + port)
}
