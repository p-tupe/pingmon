package main

import (
	"log"
	"time"

	"github.com/EMPAT94/pingmon/app"
	"github.com/EMPAT94/pingmon/config"
)

var l log.Logger = *log.Default()

func main() {
	l.Print("[INFO] Pingmon started!")

	c := config.New()
	t := *time.NewTicker(c.Interval)

	for range t.C {
		for _, site := range c.Sites {
			go app.Ping(site)
		}
	}

	l.Print("[INFO] Pingmon closed!")
}
