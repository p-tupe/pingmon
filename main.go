package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
)

type Site struct {
	URL string `json:"url"`
}

type Mailer struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	From     string `json:"from"`
}

type Config struct {
	Sites        []Site   `json:"sites"`
	Ntfy         string   `json:"ntfy,omitempty"`
	SlackWebhook string   `json:"slackWebhook,omitempty"`
	MailTo       []string `json:"mailTo,omitempty"`
	Mailer       *Mailer  `json:"mailer,omitempty"`
}

var config Config

func main() {
	// TODO: Learn how to parse flags, allow passing -c /path/to/config
	// configFile, err := os.ReadFile("/etc/pingmon/config.json")

	configFile, err := os.ReadFile("./config.json")
	Check(err)

	err = json.Unmarshal(configFile, &config)
	Check(err)

	fmt.Println(config)

	wg := &sync.WaitGroup{}
	for _, site := range config.Sites {
		wg.Add(1)
		go (func() {
			defer wg.Done()
			timer := time.NewTicker(5 * time.Second)
			for {
				<-timer.C
				ping(site.URL)
			}
		})()
	}

	wg.Wait()
}

func ping(url string) {
	resp, err := http.Get(url)

	if err != nil {
		msg := "Error while pinging site " + url
		fmt.Println(err)
		// slog.Log(resp.Request.Context(), slog.LevelError, msg)
		notify(msg)
	}

	if resp.StatusCode > 399 {
		notify("Ping failed for site: " + url)
		// slog.Log(resp.Request.Context(), slog.LevelError, "Health check for  "+url+" failed!")
	} else {
		notify("Ping succeeded for site: " + url)
		// slog.Log(resp.Request.Context(), slog.LevelInfo, "Health check for  "+url+" succeeded!")
	}
}

func notify(msg string) {
	if config.Ntfy != "" {
		_, err := http.Post(config.Ntfy, "application/text", strings.NewReader(msg))
		Check(err)
	}

	// Add slack, email config here!
}
