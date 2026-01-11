package web

import (
	"html/template"
	"io"
	"net/http"
)

func ConfigPage(tmpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tmpl.ExecuteTemplate(w, "config.html", "")
	}
}

func UpdateConfig() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "OK")
	}
}
