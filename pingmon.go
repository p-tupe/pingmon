package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	i "github.com/p-tupe/pingmon/internal"
)

func main() {
	cfg, err := i.NewConfig("./config.json")
	if err != nil {
		log.Fatalln(err)
	}

	ctx := context.Background()
	ctx, cancelJobs := context.WithCancel(ctx)

	go i.InitWriteChan(ctx)

	for _, site := range cfg.Sites {
		job, err := i.NewPingJob(site)
		if err != nil {
			log.Fatalln("Error while creating new ping job:", err.Error())
		}

		go job.Start(ctx)
	}

	quitChannel := make(chan os.Signal, 1)
	signal.Notify(quitChannel, syscall.SIGINT, syscall.SIGTERM)
	<-quitChannel
	cancelJobs()
	log.Println("Shutdown gracefully")
}
