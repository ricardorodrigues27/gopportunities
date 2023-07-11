package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func initializeRoutes(r *gin.Engine) {
	v1 := r.Group("/api/v1")
	{
		v1.GET("/openings", listOpeningsEndpoint)
		v1.GET("/openings/:id", defaultEndpoint)
		v1.POST("/openings", defaultEndpoint)
		v1.PUT("/openings/:id", defaultEndpoint)
		v1.DELETE("/openings/:id", defaultEndpoint)

	}
}

func listOpeningsEndpoint(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "listOpenings",
	})
}

func defaultEndpoint(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "defaultMessage",
	})
}
