package web

import (
	"io"
	"log"
	"net/http"
)

func SitePage() http.HandlerFunc {
	siteHTML, err := publicFS.ReadFile("public/site.html")
	if err != nil {
		log.Fatalln("Error while reading site.html, ", err)
	}
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write(siteHTML)
	}
}

func CreateSite() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "OK")
	}
}

func ReadSite() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "OK")
	}
}

func UpdateSite() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "OK")
	}
}

func DeleteSite() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "OK")
	}
}
