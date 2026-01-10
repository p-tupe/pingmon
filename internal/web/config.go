package web

import (
	"io"
	"log"
	"net/http"
)

func ConfigPage() http.HandlerFunc {
	configHTML, err := publicFS.ReadFile("public/config.html")
	if err != nil {
		log.Fatalln("Error while reading site.html, ", err)
	}
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write(configHTML)
	}
}

func UpdateConfig() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "OK")
	}
}
