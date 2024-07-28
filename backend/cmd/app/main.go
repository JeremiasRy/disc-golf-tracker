package main

import (
	"disc-golf-tracker/backend/pkg/models"
	"disc-golf-tracker/backend/pkg/repository"
	"log"

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

	courseRepository := repository.NewRepository[models.Course](db)
	roundRepository := repository.NewRepository[models.Round](db)
	holeRepository := repository.NewRepository[models.Hole](db)
	scoreRepository := repository.NewRepository[models.Score](db)
	scoreCardRepository := repository.NewRepository[models.ScoreCard](db)
	userRepository := repository.NewRepository[models.User](db)

	newCourse := models.Course{Name: "JAu!"}
	courseRepository.Create(&newCourse)

	newHole := models.Hole{NthHole: 1, CourseID: newCourse.ID, Par: 4}
	holeRepository.Create(&newHole)

	newUser := models.User{Name: "Jeremu"}
	userRepository.Create(&newUser)

	newUser2 := models.User{Name: "Jeremu2"}
	userRepository.Create(&newUser2)

	newRound := models.Round{CourseID: newCourse.ID}

	roundRepository.Create(&newRound)

	newScoreCard := models.ScoreCard{UserID: newUser.ID, RoundID: newRound.ID}
	newScoreCard2 := models.ScoreCard{UserID: newUser2.ID, RoundID: newRound.ID}

	scoreCardRepository.Create(&newScoreCard)
	scoreCardRepository.Create(&newScoreCard2)

	newScore := models.Score{HoleID: newHole.ID, ScorecardID: newScoreCard.ID, Strokes: 2, Penalties: 1}
	newScore2 := models.Score{HoleID: newHole.ID, ScorecardID: newScoreCard2.ID, Strokes: 2, Penalties: 1}

	scoreRepository.Create(&newScore)
	scoreRepository.Create(&newScore2)

}
