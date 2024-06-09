package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name string `json:"name"`
	Rounds []Round `gorm:"foreignKey:UserID" json:"rounds"`
}

type Round struct {
	gorm.Model
	UserID uint `json:"user_id"`
	CourseID uint `json:"course_id"`
	LayoutID uint `json:"layout_id"`
	ScorecardID int `json:"scorecard_id"`
}

type Score struct {
	gorm.Model
	HoleID uint `json:"hole_id"`
	ScorecardID uint `json:"scorecard_id"`
	Strokes uint `json:"strokes"`
}

type ScoreCard struct {
	gorm.Model
	Name string `json:"name"`
	Scores []Score `gorm:"foreignKey:ScorecardID" json:"scores"`
}