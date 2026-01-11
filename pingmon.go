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
	ctx, cancelJobs := context.WithCancel(context.Background())
	defer cancelJobs()

	configPath := flag.String("config", "./config.json", "set config file path")
	flag.Parse()

	log.Println("Reading config from", *configPath)
	config, err := i.NewConfig(*configPath)
	if err != nil {
		log.Fatalln(err)
	}

	go i.InitStore(ctx)
	go i.InitAlert(ctx)
	jobs := i.InitJobs(ctx)
	if config.Server.Enabled {
		go i.StartServer(ctx, jobs)
	}

	quitChannel := make(chan os.Signal, 1)
	signal.Notify(quitChannel, syscall.SIGINT, syscall.SIGTERM)
	<-quitChannel
	log.Println("Shut down gracefully")
}
