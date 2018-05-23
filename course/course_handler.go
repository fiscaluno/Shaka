package course

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/fiscaluno/shaka/db"
	"github.com/fiscaluno/shaka/util"
	"github.com/gorilla/mux"
)

// GetAll Courses
func GetAll(w http.ResponseWriter, r *http.Request) {
	db := db.Conn()
	defer db.Close()
	var courses []Course
	db.Find(&courses)
	util.RespondWithJSON(w, http.StatusOK, courses)
}

// FindByName find a Course by name
func FindByName(w http.ResponseWriter, r *http.Request) {
	db := db.Conn()
	defer db.Close()

	var courses []Course
	vars := mux.Vars(r)
	name := vars["name"]
	db.Find(&courses, "name = ?", name)
	if len(courses) >= 0 {
		util.RespondWithJSON(w, http.StatusOK, courses)
		return
	}

	msg := util.Message{
		Content: "Not exist this course",
		Status:  "ERROR",
		Body:    nil,
	}
	util.RespondWithJSON(w, http.StatusOK, msg)

}

// FindByID find a Course by ID
func FindByID(w http.ResponseWriter, r *http.Request) {
	db := db.Conn()
	defer db.Close()

	var course Course

	var msg util.Message

	msg = util.Message{
		Content: "Invalid ID, not a int",
		Status:  "ERROR",
		Body:    nil,
	}

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		util.RespondWithJSON(w, http.StatusOK, msg)
		return
	}
	db.Find(&course, id)
	if course.ID != 0 {
		util.RespondWithJSON(w, http.StatusOK, course)
		return
	}

	msg = util.Message{
		Content: "Not exist this course",
		Status:  "ERROR",
		Body:    nil,
	}
	util.RespondWithJSON(w, http.StatusOK, msg)

}

// Add a Course
func Add(w http.ResponseWriter, r *http.Request) {
	db := db.Conn()
	defer db.Close()

	var course Course
	var msg util.Message

	msg = util.Message{
		Content: "Invalid request payload",
		Status:  "ERROR",
		Body:    nil,
	}

	if err := json.NewDecoder(r.Body).Decode(&course); err != nil {
		util.RespondWithJSON(w, http.StatusBadRequest, msg)
		return
	}

	db.Create(&course)

	msg = util.Message{
		Content: "New course successfully added",
		Status:  "OK",
		Body:    course,
	}
	util.RespondWithJSON(w, http.StatusCreated, msg)

}

// DeleteByID find a Course by ID
func DeleteByID(w http.ResponseWriter, r *http.Request) {
	db := db.Conn()
	defer db.Close()

	var course Course
	var msg util.Message

	msg = util.Message{
		Content: "Invalid ID, not a int",
		Status:  "ERROR",
		Body:    nil,
	}

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		util.RespondWithJSON(w, http.StatusOK, msg)
		return
	}

	db.Find(&course, id)
	if course.ID != 0 {
		db.Delete(&course)
		msg = util.Message{
			Content: "Deleted course successfully",
			Status:  "OK",
			Body:    course,
		}

		util.RespondWithJSON(w, http.StatusAccepted, msg)
		return
	}

	msg = util.Message{
		Content: "Not exist this course",
		Status:  "ERROR",
		Body:    nil,
	}
	util.RespondWithJSON(w, http.StatusOK, msg)

}

// UpdateByID find a Course by ID
func UpdateByID(w http.ResponseWriter, r *http.Request) {
	db := db.Conn()
	defer db.Close()

	var course Course
	var msg util.Message

	msg = util.Message{
		Content: "Invalid ID, not a int",
		Status:  "ERROR",
		Body:    nil,
	}

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		util.RespondWithJSON(w, http.StatusOK, msg)
		return
	}

	msg = util.Message{
		Content: "Invalid request payload",
		Status:  "ERROR",
		Body:    nil,
	}

	if err := json.NewDecoder(r.Body).Decode(&course); err != nil {
		util.RespondWithJSON(w, http.StatusBadRequest, msg)
		return
	}

	if id >= 0 {
		course.ID = uint(id)
		db.Save(&course)

		msg = util.Message{
			Content: "Update course successfully ",
			Status:  "OK",
			Body:    course,
		}
		util.RespondWithJSON(w, http.StatusAccepted, msg)
		return
	}

	msg = util.Message{
		Content: "Not exist this course",
		Status:  "ERROR",
		Body:    nil,
	}
	util.RespondWithJSON(w, http.StatusOK, msg)

}
