package config

import (
	"fmt"
	"net/url"

	"gopkg.in/ini.v1"
)

type Job struct {
	URL      string
	Interval int
}

type Email struct {
	Id       string
	Password string
}

type Slack struct {
	webhook string
}

type Config struct {
	jobs  []Job
	email Email
	slack Slack
}

func ReadConfig() (*Config, error) {
	cfg := new(Config)

	raw, err := ini.Load("/tmp/config.ini")
	if err != nil {
		return nil, fmt.Errorf("Error while reading config: %v", err)
	}

	allSections := raw.Sections()
	for _, section := range allSections {
		if section.Name() == ini.DefaultSection {
			continue
		} else if section.Name() == "Email" {
			raw.Section("Email").MapTo(&cfg.email)
		} else if section.Name() == "Slack" {
			raw.Section("Slack").MapTo(&cfg.slack)
		} else {
			jobUrl := section.Name()

			if _, err := url.ParseRequestURI(jobUrl); err != nil {
				return nil, fmt.Errorf("Error while parsing URL in config: %v", err)
			}

			jobInterval := section.Key("Interval").MustInt()

			if jobInterval == 0 {
				jobInterval = 15
			}

			j := Job{
				URL:      jobUrl,
				Interval: jobInterval,
			}

			cfg.jobs = append(cfg.jobs, j)
		}
	}

	if len(cfg.jobs) == 0 {
		return nil, fmt.Errorf("Please add atleast one job to the config")

	}

	if cfg.email.Id == "" && cfg.slack.webhook == "" {
		return nil, fmt.Errorf("Please add either Email or Slack to the config")
	}

	err = raw.StrictMapTo(cfg)
	if err != nil {
		return nil, fmt.Errorf("Error while reading config: %v", err)
	}

	return cfg, nil
}
