package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	//initialize Gin
	router := gin.Default()

	//create route with method GET
	router.GET("/", func(c *gin.Context) {

		//return response JSON
		c.JSON(200, gin.H{
			"message": "Hello World!",
		})
	})

	//start server with port 3000
	router.Run(":3099")
}