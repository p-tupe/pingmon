package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"

	i "github.com/p-tupe/pingmon/internal"
)

func main() {
	ctx := context.Background()
	ctx, cancelJobs := context.WithCancel(ctx)

	configPath := flag.String("config", "./config.json", "set config file path")
	flag.Parse()

	// Read Config
	log.Println("Reading config from", *configPath)
	config, err := i.NewConfig(*configPath)
	if err != nil {
		log.Fatalln(err)
	}

	// Enable services
	go i.InitStore(ctx)
	go i.InitAlert(ctx)
	go i.InitJobs(ctx)
	if config.Server.Enabled {
		go i.StartServer(ctx)
	}

	// Handle Shutdown
	quitChannel := make(chan os.Signal, 1)
	signal.Notify(quitChannel, syscall.SIGINT, syscall.SIGTERM)
	<-quitChannel
	cancelJobs()
	log.Println("Shut down gracefully")
	os.Exit(0)
}
