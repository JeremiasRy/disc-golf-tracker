package main

import (
	"disc-golf-tracker/backend/pkg/models"
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
}