package controllers

import (
	"disc-golf-tracker/backend/pkg/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type RoundController struct {
	service *services.RoundService
}

type CreateRoundRequest struct {
	CourseID uint `bindings:"required" json:"course_id"`
}

func NewRoundController(service *services.RoundService) RoundController {
	return RoundController{service: service}
}

func (controller *RoundController) HandleGetRound(c *gin.Context) {
	roundId, err := strconv.Atoi(c.Param("roundId"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	round, err := controller.service.GetRound(uint(roundId))

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, round)
}

func (controller *RoundController) HandleCreateRound(c *gin.Context) {
	var requestBody CreateRoundRequest

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	round, err := controller.service.CreateRound(requestBody.CourseID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, round)
}
