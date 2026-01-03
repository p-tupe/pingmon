package web

import (
	"log"
	"net/http"
)

func HomePage() http.HandlerFunc {
	homeHTML, err := publicFS.ReadFile("public/home.html")
	if err != nil {
		log.Fatalln("Error while reading home.html, ", err)
	}

	return func(w http.ResponseWriter, r *http.Request) {
		w.Write(homeHTML)
	}
}
