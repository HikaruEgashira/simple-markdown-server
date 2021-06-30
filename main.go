package main

import (
	"net/http"

	"github.com/HikaruEgashira/simple-server/controller"
	"github.com/gorilla/mux"
)

func configureRouter(r *mux.Router) {
	r.HandleFunc("/", controller.IndexHandler)
	r.HandleFunc("/user", controller.UserHandler)
}

func main() {
	r := mux.NewRouter()
	configureRouter(r)
	http.ListenAndServe(":8080", r)
}
