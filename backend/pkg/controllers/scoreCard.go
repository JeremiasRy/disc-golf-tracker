package controllers

import (
	"disc-golf-tracker/backend/pkg/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ScoreCardController struct {
	service *services.ScoreCardService
}

type CreateScoreCardRequest struct {
	RoundID uint `bindings:"required" json:"round_id"`
	UserID  uint `bindings:"required" json:"user_id"`
}

func NewScoreCardController(service *services.ScoreCardService) ScoreCardController {
	return ScoreCardController{service: service}
}

func (controller *ScoreCardController) HandleCreateScoreCard(c *gin.Context) {
	var requestBody CreateScoreCardRequest
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	scoreCard, err := controller.service.CreateScoreCard(requestBody.RoundID, requestBody.UserID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, scoreCard)
}
