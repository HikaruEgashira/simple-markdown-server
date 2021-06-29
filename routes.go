package main

import (
	"net/http"
	"strconv"

	"github.com/HikaruEgashira/simple-server/lib"
	"github.com/HikaruEgashira/simple-server/usecase"
)

// シンプルな例
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	output := lib.GenerateHTML("pages/index")
	lib.Render(w, output)
}

// パラメータを利用する例
func AddHandler(w http.ResponseWriter, r *http.Request) {
	a, _ := strconv.Atoi(r.URL.Query().Get("a"))
	b, _ := strconv.Atoi(r.URL.Query().Get("b"))
	c := usecase.Add(a, b)

	output := lib.GenerateHTML("pages/add", map[string]int{"a": a, "b": b, "c": c})
	lib.Render(w, output)
}

func ReadmeHandler(w http.ResponseWriter, r *http.Request) {
	output := lib.GenerateHTML("README")
	lib.Render(w, output)
}
