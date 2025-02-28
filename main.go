package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"
)

var (
	sites = [...]string{
		"https://www.example.com",
	}
	slack_webhook = "https://hooks.slack.com/services/whatever/"
	interval      = 30
)

var l log.Logger = *log.Default()

func main() {
	l.Print("[INFO] Pingmon started!")
	t := *time.NewTicker(time.Duration(interval * int(time.Second)))

	for _ = range t.C {
		for _, site := range sites {
			go ping(site)
		}
	}

	l.Print("[INFO] Pingmon closed!")
}

func ping(site string) {
	l.Printf("[INFO] Pinging %s", site)

	res, err := http.Get(site)
	if err != nil {
		l.Print("[ERROR] " + err.Error())
		return
	}

	l.Printf("[INFO] Status for %s %v", site, res.StatusCode)

	ok := (res.StatusCode > 199 && res.StatusCode < 299) || (res.StatusCode >= 400 && res.StatusCode < 499)

	if !ok {
		l.Printf("[INFO] Generating alert for %s", site)
		alert(site, res.StatusCode)
	}
}

func alert(site string, statusCode int) {
	v, err := json.Marshal(map[string]string{
		"text": "Could no reach " + site +
			", received status code " + strconv.Itoa(statusCode) +
			". It may be down!",
	})

	if err != nil {
		l.Print("[ERROR] " + err.Error())
		return
	}

	_, err = http.Post(slack_webhook, "application/json", bytes.NewBuffer(v))
	if err != nil {
		l.Print("[ERROR] " + err.Error())
	}
}
