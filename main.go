package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	configureRouter(r)

	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}

func configureRouter(r *mux.Router) {
	r.HandleFunc("/", IndexHandler)
	r.HandleFunc("/add", AddHandler)
	r.HandleFunc("/readme", ReadmeHandler)
}
