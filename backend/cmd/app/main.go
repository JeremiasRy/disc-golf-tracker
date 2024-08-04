package main

import (
	"disc-golf-tracker/backend/pkg/controllers"
	"disc-golf-tracker/backend/pkg/models"
	"disc-golf-tracker/backend/pkg/repositories"
	"disc-golf-tracker/backend/pkg/services"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("disc_golf_tracker.db"))

	if err != nil {
		log.Fatal("Failed to open database connection")
	}

	err = db.AutoMigrate(&models.User{}, &models.Round{}, &models.Score{}, &models.Score{}, &models.ScoreCard{}, &models.Course{}, &models.Hole{})

	if err != nil {
		log.Fatal("Migration failed!")
	}

	courseRepo := repositories.NewRepository[models.Course](db)
	courseService := services.NewCourseService(&courseRepo)
	courseController := controllers.NewCourseController(&courseService)

	r := gin.Default()

	r.POST("/courses", courseController.HandleCreateCourse)
	r.GET("/courses", courseController.HandleGetCourses)
	r.GET("/courses/:courseId", courseController.HandleGetCourseById)
	r.Run(":8800")
}
