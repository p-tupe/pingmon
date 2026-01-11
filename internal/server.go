package internal

import (
	"context"
	"log"
	"net"
	"net/http"
)

func StartServer(ctx context.Context, jobs []*Ping) {
	log.Println("Starting server on ", cfg.Server.Addr)

	mux := http.NewServeMux()
	for path, handler := range InitRoutes(jobs) {
		mux.HandleFunc(path, handler)
	}

	server := &http.Server{
		Addr:        cfg.Server.Addr,
		BaseContext: func(l net.Listener) context.Context { return ctx },
		Handler:     mux,
	}

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Fatalln("Error while starting server:", err.Error())
		}
	}()

	<-ctx.Done()
	err := server.Shutdown(ctx)
	if err != nil {
		log.Printf("graceful shutdown failed: %v, forcing close\n", err)
		err = server.Close()
		if err != nil {
			log.Fatalf("server close failed: %v\n", err)
		}
	}
}
