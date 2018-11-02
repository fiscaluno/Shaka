package course

import "github.com/fiscaluno/pandorabox/db"

// Migrate migration entity DB
func Migrate() {
	db := db.Conn()
	defer db.Close()

	course := new(Course)

	course.Name = "Information Systems"
	course.Type = "Bachelor"
	course.Period = "nightly"
	course.Semester = 8
	course.MonthlyPaymentValue = 1000
	course.AverageRating = 5
	course.RatedByCount = 1

	// Migrate the schema
	db.AutoMigrate(&course)

	// Create
	db.Create(&course)

	// Read
	// var course Entity
	db.Find(&course)

	// Update - update course's Name to Information Systems
	db.Model(&course).Update("Name", "Information Systems")

	// Delete - delete course
	//db.Delete(&course)
}
