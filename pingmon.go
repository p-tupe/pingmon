package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"

	i "github.com/p-tupe/pingmon/internal"
	s "github.com/p-tupe/pingmon/internal/server"
)

func main() {
	ctx := context.Background()
	ctx, cancelJobs := context.WithCancel(ctx)

	startServer := flag.Bool("server", false, "start pingmon server")
	configPath := flag.String("config", "./config.json", "set config file path")
	flag.Parse()

	log.Println("Reading config from", *configPath)
	_, err := i.NewConfig(*configPath)
	if err != nil {
		log.Fatalln(err)
	}
	if *startServer {
		go s.Start(ctx)
	}
	go i.InitWriteChan(ctx)
	go i.InitAlertChan(ctx)
	i.InitJobs(ctx)

	quitChannel := make(chan os.Signal, 1)
	signal.Notify(quitChannel, syscall.SIGINT, syscall.SIGTERM)
	<-quitChannel
	cancelJobs()
	log.Println("Shutdown gracefully")
	os.Exit(0)
}
