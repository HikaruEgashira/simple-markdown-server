package lib

import (
	"html/template"
	"io/ioutil"
	"net/http"

	"github.com/cbroglie/mustache"
	"github.com/russross/blackfriday"
)

func BuildHTML(path string, context ...interface{}) string {
	md, _ := ioutil.ReadFile(path + ".md")
	mustacheHtml := string(blackfriday.Run(md))
	html, _ := mustache.Render(mustacheHtml, context...)
	return html
}

func Render(w http.ResponseWriter, data string) {
	tmpl := template.Must(template.ParseFiles("template/md.html"))
	tmpl.ExecuteTemplate(w, "md", template.HTML(data))
}
