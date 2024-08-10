package main

import (
	"disc-golf-tracker/backend/pkg/controllers"
	"disc-golf-tracker/backend/pkg/middleware"
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
	db.Exec("PRAGMA foreign_keys = ON")
	err = db.AutoMigrate(&models.User{}, &models.Round{}, &models.Score{}, &models.Score{}, &models.ScoreCard{}, &models.Course{}, &models.Hole{})

	if err != nil {
		log.Fatal("Migration failed!")
	}

	r := gin.Default()

	r.Use(middleware.CORSMiddleware())

	// COURSES
	courseRepo := repositories.NewRepository[models.Course](db)
	courseService := services.NewCourseService(&courseRepo)
	courseController := controllers.NewCourseController(&courseService)

	r.POST("/courses", courseController.HandleCreateCourse)
	r.GET("/courses", courseController.HandleGetCourses)
	r.GET("/courses/:courseId", courseController.HandleGetCourseById)

	// HOLES
	holeRepo := repositories.NewRepository[models.Hole](db)
	holeService := services.NewHoleService(&holeRepo)
	holeController := controllers.NewHoleController(&holeService)

	r.POST("/holes", holeController.HandleCreateHole)
	r.GET("/holes/:holeId", holeController.HandleGetHoleById)

	// USERS
	userRepo := repositories.NewRepository[models.User](db)
	userService := services.NewUserService(&userRepo)
	userController := controllers.NewUserController(&userService)

	r.POST("/users", userController.HandleCreateUser)
	r.GET("/users", userController.HandleGetAllUsers)
	r.GET("/users/:userId", userController.HandleGetUserById)
	r.PATCH("/users/:userId", userController.HandleUpdateUser)

	// SCORES
	scoreRepo := repositories.NewRepository[models.Score](db)
	scoreService := services.NewScoreService(&scoreRepo)
	scoreController := controllers.NewScoreController(&scoreService)

	r.POST("/scores", scoreController.HandleCreateScore)
	r.PATCH("/scores/:scoreId", scoreController.HandleEditScore)

	// SCORECARDS
	scoreCardRepo := repositories.NewRepository[models.ScoreCard](db)
	scoreCardService := services.NewScoreCardService(&scoreCardRepo)
	scoreCardController := controllers.NewScoreCardController(&scoreCardService)

	r.POST("/scorecards", scoreCardController.HandleCreateScoreCard)

	// ROUNDS
	roundRepo := repositories.NewRepository[models.Round](db)
	roundService := services.NewRoundService(&roundRepo)
	roundController := controllers.NewRoundController(&roundService)

	r.POST("/rounds", roundController.HandleCreateRound)
	r.GET("/rounds/:roundId", roundController.HandleGetRound)

	r.Run(":8800")
}
