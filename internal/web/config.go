package web

import (
	"io"
	"net/http"
)

func ConfigPage() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "OK")
	}
}

func UpdateConfig() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "OK")
	}
}
