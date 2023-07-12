package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ricardorodrigues27/gopportunities.git/schemas"
)

func ListOpeningsHandler(c *gin.Context) {
	openings := []schemas.Opening{}

	if err := db.Find(&openings).Error; err != nil {
		sendError(c, http.StatusInternalServerError, "error listing openings")
		return
	}

	sendSuccess(c, "ListOpenings", openings)
}
