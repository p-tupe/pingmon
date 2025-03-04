package main

import (
	"fmt"
	"log"

	"time"

	"github.com/EMPAT94/pingmon/internal/app"
	"github.com/EMPAT94/pingmon/internal/cmd/help"
	"github.com/EMPAT94/pingmon/internal/config"
	"github.com/EMPAT94/pingmon/internal/flag"
)

var l log.Logger = *log.Default()

func main() {
	l.Print("[INFO] Pingmon started!")

	cmd, err := flag.Parse()

	if err != nil {
		fmt.Println(err)
		return
	}

	switch cmd {
	case "help":
		help.Show()
	}

	c := config.New()
	t := *time.NewTicker(c.Interval)

	for range t.C {
		for _, site := range c.Sites {
			go app.Ping(site)
		}
	}

	l.Print("[INFO] Pingmon closed!")
}
