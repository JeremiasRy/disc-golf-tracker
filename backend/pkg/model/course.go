package model

import "gorm.io/gorm"

type Course struct {
	gorm.Model
	Name string `json:"name"`
	Layouts []Layout `json:"layouts"` // course has-many layouts
	Holes []Hole `json:"holes"` // course has-many holes
}

type Hole struct {
	gorm.Model
	Par uint `json:"par"`
	NthHole uint `json:"nth_hole"`
	CourseID uint `json:"course_id"` // foreign key
}

type Layout struct {
	gorm.Model
	CourseID uint `json:"course_id"` // foreign key
	Name string `json:"name"`
	Holes []Hole `gorm:"many2many:layout_holes"`
}

type LayoutHole struct { // mapping table for layouts
	LayoutID uint 
	HoleID uint 
	Order uint
}