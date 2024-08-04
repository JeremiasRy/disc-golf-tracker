package controllers

import (
	"disc-golf-tracker/backend/pkg/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type HoleController struct {
	service *services.HoleService
}

type CreateHoleRequest struct {
	NthHole  uint `json:"nth_hole" binding:"required"`
	Par      uint `json:"par" binding:"required"`
	CourseId uint `json:"course_id" binding:"required"`
}

func NewHoleController(service *services.HoleService) HoleController {
	return HoleController{service: service}
}

func (controller *HoleController) HandleGetHoleById(c *gin.Context) {
	holeId, err := strconv.Atoi(c.Param("courseId"))

	if err != nil || holeId < 0 {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}

	course, err := controller.service.GetHole(uint(holeId))

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, course)
}

func (controller *HoleController) HandleCreateHole(c *gin.Context) {
	var requestBody CreateHoleRequest
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hole, err := controller.service.CreateHole(requestBody.NthHole, requestBody.Par, requestBody.CourseId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, hole)
}
