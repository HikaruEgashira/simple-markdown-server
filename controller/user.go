package controller

import (
	"net/http"

	"github.com/HikaruEgashira/simple-server/lib"
	"github.com/HikaruEgashira/simple-server/usecase"
)

func UserHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")

	text := usecase.Hello(name)

	output := lib.BuildHTML("pages/user", map[string]string{"text": text})
	lib.Render(w, output)
}
