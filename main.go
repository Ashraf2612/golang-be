package main

import (
	"golang-BE-app/configs"
	"golang-BE-app/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	configs.ConnectDB()

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"data": "Hello world!",
		})
	})

	routes.UserRoute(router)

	router.Run("localhost:3000")
}
