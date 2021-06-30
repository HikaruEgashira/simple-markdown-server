package controller

import (
	"net/http"

	"github.com/HikaruEgashira/simple-server/lib"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	output := lib.BuildHTML("pages/index")
	lib.Render(w, output)
}
