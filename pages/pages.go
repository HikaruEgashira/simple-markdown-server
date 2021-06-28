package pages

import (
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/cbroglie/mustache"
	"github.com/russross/blackfriday"
)

func generateHTML(path string) string {
	out, _ := ioutil.ReadFile(path + ".md")
	return string(blackfriday.Run(out))
}

// シンプルな例
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	output := generateHTML("pages/index")
	tmpl := template.Must(template.ParseFiles("template/md.html"))
	if err := tmpl.ExecuteTemplate(w, "md", template.HTML(output)); err != nil {
		InternalServerErrorPage(err, w)
	}
}

// パラメータを利用する例
func AddHandler(w http.ResponseWriter, r *http.Request) {
	output := generateHTML("pages/add")
	a, _ := strconv.Atoi(r.URL.Query().Get("a"))
	b, _ := strconv.Atoi(r.URL.Query().Get("b"))
	c := a + b

	data, _ := mustache.Render(output, map[string]int{"a": a, "b": b, "c": c})
	tmpl := template.Must(template.ParseFiles("template/md.html"))
	if err := tmpl.ExecuteTemplate(w, "md", template.HTML(data)); err != nil {
		InternalServerErrorPage(err, w)
	}
}

func ReadmeHandler(w http.ResponseWriter, r *http.Request) {
	output := generateHTML("README")
	tmpl := template.Must(template.ParseFiles("template/md.html"))
	if err := tmpl.ExecuteTemplate(w, "md", template.HTML(output)); err != nil {
		InternalServerErrorPage(err, w)
	}
}

func InternalServerErrorPage(err error, w http.ResponseWriter) {
	log.Panic(err)
	text := `# Error 500
Internal server error`
	output := template.HTML(string(blackfriday.Run([]byte(text))))
	tmpl := template.Must(template.ParseFiles("template/md.html"))
	tmpl.ExecuteTemplate(w, "md", template.HTML(output))
}
