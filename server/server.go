package server

import (
	"log"
	"net/http"

	"github.com/fiscaluno/pandorabox/mu"
	"github.com/fiscaluno/shaka/course"
	"github.com/fiscaluno/shaka/logs"
	"github.com/fiscaluno/shaka/util"
	"github.com/gorilla/mux"
)

var name string

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Ola, Seja bem vindo ao Shaka @" + name + " !!"))
}

// Listen init a http server
func Listen() {
	port := util.GetOSEnvironment("PORT", "5002")

	name = util.GetOSEnvironment("NAME", "JC")

	r := mux.NewRouter()
	r.Use(logs.LoggingMiddleware)
	r.Use(mu.AuthMiddleware)

	course.SetRoutes(r.PathPrefix("/courses").Subrouter())

	r.HandleFunc("/", handler)
	http.Handle("/", r)

	log.Println("Listen on port: " + port)
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatal(err)
	}
}
