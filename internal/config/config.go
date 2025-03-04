package config

import "time"

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
