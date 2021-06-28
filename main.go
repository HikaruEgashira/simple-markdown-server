package main

import (
	"log"
	"net/http"

	"github.com/HikaruEgashira/simple-server/pages"
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
	r.HandleFunc("/", pages.IndexHandler)
	r.HandleFunc("/add", pages.AddHandler)
	r.HandleFunc("/readme", pages.ReadmeHandler)
}
