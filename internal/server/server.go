package server

import (
	"context"
	"io"
	"log"
	"net"
	"net/http"
)

func Start(ctx context.Context) {
	log.Println("Starting server on :8080")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "OK")
	})

	server := &http.Server{
		Addr:        ":8080",
		BaseContext: func(l net.Listener) context.Context { return ctx },
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatalln("Error while starting server:", err.Error())
	}
}
