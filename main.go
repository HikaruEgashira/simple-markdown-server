package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	configureRouter(r)
	http.ListenAndServe(":8080", r)
}

func configureRouter(r *mux.Router) {
	r.HandleFunc("/", IndexHandler)
	r.HandleFunc("/add", AddHandler)
	r.HandleFunc("/user", UserHandler)
	r.HandleFunc("/readme", ReadmeHandler)
}
