package course

import "github.com/gorilla/mux"

// SetRoutes add routes from course
func SetRoutes(r *mux.Router) {
	r.HandleFunc("", Add).Methods("POST")
	r.HandleFunc("", GetAll).Methods("GET")
	r.HandleFunc("/names/{name}", FindByName).Methods("GET")
	r.HandleFunc("/{id}", FindByID).Methods("GET")
	r.HandleFunc("/{id}", DeleteByID).Methods("DELETE")
	r.HandleFunc("/{id}", UpdateByID).Methods("PUT")
}
