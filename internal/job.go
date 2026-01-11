package internal

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

type Ping struct {
	OK       bool
	LastPing int64 // Unix second
	URL      string
	Interval time.Duration
}

func InitJobs(ctx context.Context) []*Ping {
	var jobs []*Ping
	for _, site := range cfg.Sites {
		job, err := NewPingJob(site)
		if err != nil {
			log.Fatalln("Error while creating new ping job:", err.Error())
		}

		go job.start(ctx)
		jobs = append(jobs, job)
	}

	return jobs
}

func NewPingJob(site Site) (*Ping, error) {
	ping := &Ping{URL: site.URL, Interval: site.Interval}

	if site.Interval < 1 {
		ping.Interval = DEFAULT_INTERVAL
	}

	return ping, nil
}

func (ping *Ping) start(ctx context.Context) {
	log.Printf("Starting job for URL %s every %s\n", ping.URL, ping.Interval*time.Second)
	ticker := time.Tick(ping.Interval * time.Second)

	for {
		select {
		case <-ctx.Done():
			return

		case <-ticker:
			log.Println("Pinging", ping.URL)
			checkSite(ping)
		}
	}
}

func checkSite(ping *Ping) {
	ping.LastPing = time.Now().Unix()
	ping.OK = true
	msg := ""

	// TODO: Add deadline/timeout etc
	resp, err := http.Get(ping.URL)

	if err != nil {
		ping.OK = false
		msg = fmt.Sprintf("Error while pinging site %s, error: %v", ping.URL, err.Error())
	} else if resp.StatusCode > 399 {
		ping.OK = false
		msg = fmt.Sprintf("Error while pinging site %s, status: %s", ping.URL, resp.Status)
	} else {
		msg = fmt.Sprintf("URL %s Status %s", ping.URL, resp.Status)
	}

	if !ping.OK {
		SendAlert(msg)
	}

	log.Println(msg)
	WriteToStore(ping)
}
