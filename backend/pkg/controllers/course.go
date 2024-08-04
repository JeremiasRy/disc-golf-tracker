package controllers

import (
	"disc-golf-tracker/backend/pkg/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CourseController struct {
	service *services.CourseService
}

type CreateCourseRequest struct {
	CourseName string `json:"course_name" binding:"required,min=3"`
}

type UpdateCourseRequest struct {
	CourseName string `json:"course_name" binding:"required,min=3"`
}

func NewCourseController(service *services.CourseService) CourseController {
	return CourseController{service: service}
}

func (controller *CourseController) HandleGetCourses(c *gin.Context) {
	courses, err := controller.service.GetAllCourses()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, courses)
}

func (controller *CourseController) HandleGetCourseById(c *gin.Context) {
	courseId, err := strconv.Atoi(c.Param("courseId"))

	if err != nil || courseId < 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	course, err := controller.service.GetCourse(uint(courseId))

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Course not found"})
		return
	}

	c.JSON(http.StatusOK, course)
}

func (controller *CourseController) HandleCreateCourse(c *gin.Context) {
	var requestBody CreateCourseRequest

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	course, err := controller.service.CreateCourse(requestBody.CourseName)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, course)
}

func (controller *CourseController) HandleUpdateCourse(c *gin.Context) {
	courseId, err := strconv.Atoi("courseId")

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var requestBody UpdateCourseRequest

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = controller.service.EditCourseName(requestBody.CourseName, uint(courseId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
