package models

import "gorm.io/gorm"

type Course struct {
	gorm.Model
	Name   string  `json:"name"`
	Holes  []Hole  `gorm:"foreignKey:CourseID" json:"holes"`
	Rounds []Round `gorm:"foreignKey:CourseID" json:"rounds"`
}

type Hole struct {
	gorm.Model
	Par      uint    `json:"par"`
	NthHole  uint    `gorm:"uniqueIndex" json:"nth_hole"`
	CourseID uint    `json:"course_id"`
	Scores   []Score `gorm:"foreignKey:HoleID" json:"scores"`
}

type Round struct {
	gorm.Model
	CourseID   uint
	ScoreCards []ScoreCard `gorm:"foreignKey:RoundID" json:"cards"`
}

type ScoreCard struct {
	gorm.Model
	RoundID uint
	UserID  uint
	Scores  []Score `gorm:"foreignKey:ScorecardID" json:"scores"`
}

type Score struct {
	gorm.Model
	HoleID      uint `json:"hole_id"`
	ScorecardID uint `json:"scorecard_id"`
	Strokes     uint `json:"strokes"`
	Penalties   uint `json:"penalties"`
}

type User struct {
	gorm.Model
	Name       string      `json:"name"`
	ScoreCards []ScoreCard `gorm:"foreignKey:UserID" json:"score_card"`
}
