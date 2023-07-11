package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ricardorodrigues27/gopportunities.git/schemas"
)

func CreateOpeningHandler(c *gin.Context) {
	request := CreateOpeningRequest{}

	c.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		sendError(c, http.StatusBadRequest, err.Error())
		return
	}

	opening := schemas.Opening{
		Role:     request.Role,
		Company:  request.Company,
		Location: request.Location,
		Remote:   *request.Remote,
		Link:     request.Link,
		Salary:   request.Salary,
		Currency: request.Currency,
	}

	if err := db.Create(&opening).Error; err != nil {
		logger.Errorf("error creating opening: %v", err.Error())
		sendError(c, http.StatusInternalServerError, err.Error())
		return
	}

	sendSuccess(c, "CreateOpening", opening)
}
