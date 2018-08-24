package course

import (
	"github.com/fiscaluno/pandorabox/db"
)

// Course is a Entity
type Course struct {
	Name          string  `json:"course_name"`
	Type          string  `json:"course_type"`
	AverageRating float64 `json:"course_average_rating"`
	RatedByCount  int     `json:"course_rated_by_count"`
}

// Entity is a course
type Entity struct {
	Course
	db.CommonModelFields
}

// Entitys is Entity slice
type Entitys []Entity

// GetAll Entitys
func GetAll() Entitys {
	db := db.Conn()
	defer db.Close()
	var entitys Entitys
	db.Find(&entitys)
	return entitys
}

// Save a Entity
func (entity Entity) Save() (Entity, error) {
	db := db.Conn()
	defer db.Close()

	db.Create(&entity)

	return entity, nil
}

// GetByID a Entity
func GetByID(id int) Entity {
	db := db.Conn()
	defer db.Close()

	var entity Entity

	db.Find(&entity, id)

	return entity
}

// GetByQuery a Entity
func GetByQuery(query string, value interface{}) Entitys {
	db := db.Conn()
	defer db.Close()

	var entitys Entitys

	db.Find(&entitys, query, value)
	return entitys
}
