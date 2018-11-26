package course

import (
	"time"

	"github.com/fiscaluno/pandorabox/db"
)

// Course is a Course
type Course struct {
	ID                  uint       `gorm:"primary_key" json:"course_id"`
	Name                string     `json:"course_name"`
	Type                string     `json:"course_type"`
	AverageRating       float64    `json:"course_average_rating"`
	RatedByCount        int        `json:"course_rated_by_count"`
	Period              string     `json:"period"`
	Semester            int        `json:"semester"`
	MonthlyPaymentValue int        `json:"monthly_payment_value"`
	CreatedAt           time.Time  `json:"created_at"`
	UpdatedAt           time.Time  `json:"updated_at"`
	DeletedAt           *time.Time `json:"deleted_at"`
	InstitutionID       int        `json:"institution_id"`
	InstitutionName     string     `json:"institution_name"`
	InstitutionImageURL string     `json:"institution_image_url"`
}

// TableName for course
func (Course) TableName() string {
	return "course"
}

// GetAll Courses
func GetAll() []Course {
	db := db.Conn()
	defer db.Close()
	var entitys []Course
	db.Find(&entitys)
	return entitys
}

// Save a Course
func (entity Course) Save() (Course, error) {
	db := db.Conn()
	defer db.Close()

	db.Create(&entity)

	return entity, nil
}

// GetByID a Course
func GetByID(id int) Course {
	db := db.Conn()
	defer db.Close()

	var entity Course

	db.Find(&entity, id)

	return entity
}

// GetByQuery a Course
func GetByQuery(query string, value interface{}) []Course {
	db := db.Conn()
	defer db.Close()

	var entitys []Course

	db.Find(&entitys, query, value)
	return entitys
}
