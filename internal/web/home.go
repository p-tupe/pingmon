package web

import (
	"html/template"
	"net/http"
)

func HomePage(tmpl *template.Template, data any) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tmpl.ExecuteTemplate(w, "home.html", data)
	}
}
