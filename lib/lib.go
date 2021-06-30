package lib

import (
	"html/template"
	"io"
	"io/ioutil"

	"github.com/cbroglie/mustache"
	"github.com/russross/blackfriday"
)

func Render(w io.Writer, path string, context ...interface{}) {
	md, _ := ioutil.ReadFile(path + ".md")
	mustacheHtml := string(blackfriday.Run(md))
	html, _ := mustache.Render(mustacheHtml, context...)

	tmpl := template.Must(template.ParseFiles("template/md.html"))
	tmpl.ExecuteTemplate(w, "md", template.HTML(html))
}
