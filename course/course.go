package course

import "github.com/jinzhu/gorm"

// Course is a Human
type Course struct {
	gorm.Model
	Name string `json:"name"`
}

// Courses is a slice for Course
// var Courses []Course
