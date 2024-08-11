package controllers

import (
	"disc-golf-tracker/backend/pkg/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ScoreController struct {
	service *services.ScoreService
}

type CreateScoreRequest struct {
	HoleID      uint `bindings:"required" json:"hole_id"`
	ScoreCardID uint `bindings:"required" json:"score_card_id"`
}

type UpdateScoreRequest struct {
	NewStrokes   uint `json:"new_strokes"`
	NewPenalties uint `json:"new_penalties"`
}

func NewScoreController(service *services.ScoreService) ScoreController {
	return ScoreController{service: service}
}

func (controller *ScoreController) HandleCreateScore(c *gin.Context) {
	var requestBody CreateScoreRequest
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	score, err := controller.service.CreateScore(requestBody.HoleID, requestBody.ScoreCardID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusAccepted, score)
}

func (controller *ScoreController) HandleEditScore(c *gin.Context) {
	scoreId, err := strconv.Atoi(c.Param("scoreId"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	requestBody := UpdateScoreRequest{NewStrokes: 0, NewPenalties: 0}

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	score, err := controller.service.UpdateScore(uint(scoreId), requestBody.NewStrokes, requestBody.NewPenalties)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, score)
}
