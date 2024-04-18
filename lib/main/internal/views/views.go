package views

import (
	"html/template"
	"io"
	"net/http"
)

type Templates struct {
	templates *template.Template
}

func (t *Templates) Render(w io.Writer, name string, data interface{}) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func NewTemplate() *Templates {
	return &Templates{
		templates: template.Must(template.ParseGlob("internal/views/*.html")),
	}
}

func RenderHTML(w http.ResponseWriter, status int, templateFile string, data interface{}) {
	tmpl := template.Must(template.ParseFiles(templateFile))
	tmpl.Execute(w, data)
}
