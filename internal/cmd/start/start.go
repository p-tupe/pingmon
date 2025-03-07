package start

import (
	"log"
	"time"

	"github.com/EMPAT94/pingmon/internal/app"
	"github.com/EMPAT94/pingmon/internal/config"
)

func Start() {
	var l log.Logger = *log.Default()

	l.Print("[INFO] Pingmon started!")

	c := config.Parse()
	t := *time.NewTicker(c.Interval)

	for range t.C {
		for _, site := range c.Sites {
			go app.Ping(site)
		}
	}

	l.Print("[INFO] Pingmon closed!")
}
