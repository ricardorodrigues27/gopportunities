package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ricardorodrigues27/gopportunities.git/schemas"
)

func UpdateOpeningHandler(c *gin.Context) {
	id := c.Param("id")

	if id == "" {
		sendError(c, http.StatusBadRequest, errParamIsRequired("id", "pathParameter").Error())
		return
	}

	opening := schemas.Opening{}

	if err := db.First(&opening, id).Error; err != nil {
		sendError(c, http.StatusNotFound, fmt.Sprintf("opening with id: %s not found", id))
		return
	}

	request := UpdateOpeningRequest{}

	c.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		sendError(c, http.StatusBadRequest, err.Error())
		return
	}

	logger.Debug(request.Map())

	if err := db.Model(&opening).Updates(request.Map()).Error; err != nil {
		logger.Errorf("error updating opening: %v", err)
		return
	}

	sendSuccess(c, "UpdateOpening", opening)
}
