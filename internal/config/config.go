package config

import "time"

var Cmds = map[string]string{
	"help":    "Prints this help",
	"init":    "Initialize pingmon site monitoring service",
	"test":    "Test network and notifications",
	"status":  "Show latest pingmon results",
	"log":     "Show latest pingmon logs",
	"start":   "Start pingmon service, if not already running",
	"restart": "Gracefully stop pingmon service",
	"stop":    "Clear all pingmon data and stop service",
	"remove":  "Clear all pingmon data and stop service",
}

type Config struct {
	Sites        []string
	SlackWebhook string
	Interval     time.Duration
}

func New() *Config {
	return &Config{
		Sites: []string{
			"https://www.example.com",
			"https://www.priteshtupe.com",
		},
		Interval:     30 * time.Minute,
		SlackWebhook: "",
	}
}
