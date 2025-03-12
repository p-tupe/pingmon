package config

import (
	"encoding/json"
	"os"
	"path/filepath"
	"time"
)

var Cmds = map[string]string{
	"help":    "Prints this help",
	"setup":   "Setup pingmon config and service files",
	"test":    "Test network and notifications",
	"status":  "Show latest pingmon results",
	"log":     "Show latest pingmon logs",
	"start":   "Start pingmon service, if not already running",
	"restart": "Gracefully stop pingmon service",
	"stop":    "Clear all pingmon data and stop service",
	"remove":  "Clear all pingmon data and stop service",
}

type Site struct {
	URL          string
	Interval     time.Duration `json:",omitempty"`
	UpStatus     int           `json:",omitempty"`
	DownStatus   int           `json:",omitempty"`
	EmailTo      []string      `json:",omitempty"`
	SlackWebhook string        `json:",omitempty"`
	Mailer       *Mailer       `json:",omitempty"`
}

type Mailer struct {
	Host     string
	Port     int
	Username string
	Password string
	From     string
}

type Config struct {
	Sites        []Site
	Interval     time.Duration
	SlackWebhook string   `json:",omitempty"`
	EmailTo      []string `json:",omitempty"`
	Mailer       *Mailer  `json:",omitempty"`
}

func Parse() (*Config, error) {
	config := Config{}

	userConfigDir, err := os.UserConfigDir()
	if err != nil {
		return nil, err
	}

	pingmonConfigDir := filepath.Join(userConfigDir, "pingmon")

	err = os.MkdirAll(pingmonConfigDir, os.ModeDir.Perm())
	if err != nil {
		return nil, err
	}

	configFilePath := filepath.Join(pingmonConfigDir, "/config.json")

	raw, err := os.ReadFile(configFilePath)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(raw, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
