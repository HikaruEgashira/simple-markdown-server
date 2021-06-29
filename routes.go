package main

import (
	"html/template"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/HikaruEgashira/simple-server/usecase"
	"github.com/cbroglie/mustache"
	"github.com/russross/blackfriday"
)

func generateHTML(path string) string {
	out, _ := ioutil.ReadFile(path + ".md")
	return string(blackfriday.Run(out))
}

func render(w http.ResponseWriter, data string) {
	tmpl := template.Must(template.ParseFiles("template/md.html"))
	tmpl.ExecuteTemplate(w, "md", template.HTML(data))
}

// シンプルな例
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	output := generateHTML("pages/index")
	render(w, output)
}

// パラメータを利用する例
func AddHandler(w http.ResponseWriter, r *http.Request) {
	output := generateHTML("pages/add")
	a, _ := strconv.Atoi(r.URL.Query().Get("a"))
	b, _ := strconv.Atoi(r.URL.Query().Get("b"))

	c := usecase.Add(a, b)

	data, _ := mustache.Render(output, map[string]int{"a": a, "b": b, "c": c})

	render(w, data)
}

func ReadmeHandler(w http.ResponseWriter, r *http.Request) {
	output := generateHTML("README")
	render(w, output)
}
