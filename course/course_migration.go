package course

import "github.com/fiscaluno/shaka/db"

// Migrate migration Course BD
func Migrate() {
	db := db.Conn()
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&Course{})

	// Create
	db.Create(&Course{Name: "J"})

	// Read
	var course Course
	db.First(&course, 1)               // find course with id 1
	db.First(&course, "name = ?", "J") // find course with code l1212

	// Update - update course's price to 2000
	db.Model(&course).Update("Name", "JC")

	// Delete - delete course
	// db.Delete(&course)
}
