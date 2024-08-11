package models

import "gorm.io/gorm"

type Course struct {
	gorm.Model
	Name   string  `json:"name"`
	Holes  []Hole  `gorm:"foreignKey:CourseID;constraint:OnDelete:Cascade" json:"holes"`
	Rounds []Round `gorm:"foreignKey:CourseID;constraint:OnDelete:Cascade" json:"rounds"`
}

type Hole struct {
	gorm.Model
	Par      uint    `json:"par"`
	NthHole  uint    `gorm:"uniqueIndex" json:"nth_hole"`
	CourseID uint    `json:"course_id"`
	Scores   []Score `gorm:"foreignKey:HoleID;constraint:OnDelete:Cascade" json:"scores"`
}

type Round struct {
	gorm.Model
	CourseID   uint
	Course     Course      `json:"course"`
	ScoreCards []ScoreCard `gorm:"foreignKey:RoundID;constraint:OnDelete:Cascade" json:"cards"`
}

type ScoreCard struct {
	gorm.Model
	RoundID uint
	UserID  uint
	User    User    `json:"user"`
	Scores  []Score `gorm:"foreignKey:ScorecardID;constraint:OnDelete:Cascade" json:"scores"`
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
	Email      string      `json:"email" gorm:"uniqueIndex"`
	ScoreCards []ScoreCard `gorm:"foreignKey:UserID;constraint:OnDelete:Cascade" json:"score_card"`
}
