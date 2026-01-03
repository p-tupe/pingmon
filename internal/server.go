package internal

import (
	"context"
	"log"
	"net"
	"net/http"

	"github.com/p-tupe/pingmon/internal/web"
)

func StartServer(ctx context.Context) {
	log.Println("Starting server on ", cfg.Server.Addr)

	mux := http.NewServeMux()
	for path, handler := range web.Routes {
		mux.HandleFunc(path, handler)
	}

	server := &http.Server{
		Addr:        cfg.Server.Addr,
		BaseContext: func(l net.Listener) context.Context { return ctx },
		Handler:     mux,
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Fatalln("Error while starting server:", err.Error())
	}
}
