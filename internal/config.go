package internal

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"
)

type Site struct {
	// Check at every x second interval
	// Defaults to 30m
	Interval time.Duration `json:"interval,omitempty"`

	// Web URL of the site
	URL string `json:"url"`
}

type Mailer struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	From     string `json:"from"`
}

type PostRequest struct {
	// URL to send the post request to
	URL string `json:"url"`

	// ContentType of the request body
	ContentType string `json:"contentType"`

	// Authorization Token Header
	// If "Bearer xxxyyzz", include "Bearer" as well
	Authorization string `json:"authorization,omitempty"`
}

type Server struct {
	// Address of server in host:port format
	// Defaults to :8080
	Addr string `json:"host,omitempty"`

	// Should start the server?
	// Defaults to true
	Enabled bool `json:"enabled,omitempty"`
}

type Config struct {
	// List of sites to ping
	Sites []Site `json:"sites"`

	// Recipient email IDs of alert messages
	MailTo []string `json:"mailTo,omitempty"`

	// Host email service to send alerts
	Mailer *Mailer `json:"mailer,omitempty"`

	// Custom webhook to send alert on
	PostRequest *PostRequest `json:"postRequest,omitempty"`

	// Path/file to save state in
	// Defaults to "./pingmon.csv"
	Store string `json:"store,omitempty"`

	// Server config
	Server Server `json:"server"`
}

var ErrNoSite = errors.New("Config must have atleast one site to monitor!")
var ErrNoAlert = errors.New("Config must have atleast one alert mechanism!")
var ErrNoMailTo = errors.New("Config must have atleast one recipient 'mailTo'!")

var cfg *Config

func NewConfig(path string) (*Config, error) {
	cfg = &Config{
		Sites:  []Site{},
		MailTo: []string{},
		Server: Server{
			Addr:    ":8080",
			Enabled: true,
		},
	}

	configFile, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("Error while reading file %v", err)
	}

	err = json.Unmarshal(configFile, &cfg)
	if err != nil {
		return nil, fmt.Errorf("Error while unmarshalling file %v", err)
	}

	if len(cfg.Sites) == 0 {
		return nil, ErrNoSite
	}

	if cfg.Mailer == nil && cfg.PostRequest == nil {
		return nil, ErrNoAlert
	}

	if cfg.Mailer != nil && len(cfg.MailTo) == 0 {
		return nil, ErrNoMailTo
	}

	if cfg.Store == "" {
		cfg.Store = DEFAULT_STORE
	}

	if cfg.PostRequest != nil && cfg.PostRequest.ContentType == "" {
		cfg.PostRequest.ContentType = DEFAULT_POSTREQUEST_CONTENTTYPE
	}

	return cfg, nil
}
