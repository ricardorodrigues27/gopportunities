package router

import "github.com/gin-gonic/gin"

func Initialize() {
	// Init Router with gin Default config
	r := gin.Default()

	// Defining routes
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Run(":8080") // listen and serve on 0.0.0.0:8080
}
