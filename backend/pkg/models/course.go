package models

import "gorm.io/gorm"

type Course struct {
	gorm.Model
	Name string `json:"name"`
	Holes []Hole `gorm:"foreignKey:CourseID" json:"holes"` // course has-many holes
}

type Hole struct {
	gorm.Model
	Par uint `json:"par"`
	NthHole uint `json:"nth_hole"`
	CourseID uint `json:"course_id"` // foreign key
}
