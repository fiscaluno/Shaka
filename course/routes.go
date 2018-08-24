package course

import (
	"github.com/gorilla/mux"
)

// SetRoutes add routes from course
func SetRoutes(r *mux.Router) {
	r.HandleFunc("", Add).Methods("POST")
	r.HandleFunc("", FindAll).Methods("GET")
	r.HandleFunc("/{id}", FindByID).Methods("GET")
	r.HandleFunc("/{id}", DeleteByID).Methods("DELETE")
	r.HandleFunc("/{id}", UpdateByID).Methods("PUT")
}
