package controllers

import (
	"disc-golf-tracker/backend/pkg/models"
	"disc-golf-tracker/backend/pkg/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	service *services.UserService
}

type CreateUserRequest struct {
	Name  string `binding:"required,min=2" json:"name"`
	Email string `binding:"required,email" json:"email"`
}

type EditUserRequest struct {
	Name string `binding:"required,min=2" json:"name"`
}

type UserResponse struct {
	Name       string
	ScoreCards []models.ScoreCard
}

func NewUserController(service *services.UserService) UserController {
	return UserController{service: service}
}

func (controller *UserController) HandleCreateUser(c *gin.Context) {
	var requestBody CreateUserRequest

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := controller.service.CreateUser(requestBody.Name, requestBody.Email)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, user)
}

// REMINDER: should only allow changing if client authenticated is the user being requested
func (controller *UserController) HandleUpdateUser(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("userId"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	var requestBody EditUserRequest

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = controller.service.EditUserName(requestBody.Name, uint(userId))

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}

func (controller *UserController) HandleGetAllUsers(c *gin.Context) {
	users, err := controller.service.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := make([]UserResponse, len(*users))

	for i, user := range *users {
		userResponse := UserResponse{Name: user.Name, ScoreCards: user.ScoreCards}
		response[i] = userResponse
	}

	c.JSON(http.StatusOK, response)
}

// REMINDER: return full object if client making the request is authenticated as the user being requested
func (controller *UserController) HandleGetUserById(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("userId"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := controller.service.GetUser(uint(userId))

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}
