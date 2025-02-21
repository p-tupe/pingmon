package config

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"strings"

	"github.com/adrg/xdg"
	"gopkg.in/ini.v1"
)

type Job struct {
	URL      string
	Interval int
}

type Mailer struct {
	Id       string
	Password string
	Server   string
}

type Slack struct {
	Webhook string
}

type Recipients struct {
	Id []string
}

type Config struct {
	Jobs       []Job
	Mailer     Mailer
	Recipients Recipients
	Slack      Slack
}

func ReadConfig(customPath string) (*Config, error) {
	var err error
	cfg := new(Config)

	configFilePath := ""

	if customPath != "" {
		_, err = os.Stat(customPath)
		if err != nil {
			return nil, fmt.Errorf("Error while reading custom config file path: %v", err)
		}
		configFilePath = customPath
	} else {
		configFilePath, err = xdg.SearchConfigFile("pingmon/config.ini")
		if err != nil {
			if strings.HasPrefix("could not locate", err.Error()) {
				configFilePath, err = xdg.ConfigFile("pingmon/config.ini")
				if err != nil {
					return nil, fmt.Errorf("Error while creating new config file: %v", err)
				}

				_, err = os.Create(configFilePath)
				if err != nil {
					return nil, fmt.Errorf("Error while creating new config file: %v", err)
				}
			} else {
				return nil, fmt.Errorf("Error while reading config file: %v", err)
			}
		}
	}

	log.Println("Loading config from:", configFilePath)

	raw, err := ini.Load(configFilePath)
	if err != nil {
		return nil, fmt.Errorf("error while reading config: %v", err)
	}

	allSections := raw.Sections()
	for _, section := range allSections {

		switch section.Name() {
		case ini.DefaultSection:
			continue

		case "Mailer":
			raw.Section("Mailer").MapTo(&cfg.Mailer)

		case "Recipients":
			raw.Section("Recipients").MapTo(&cfg.Recipients)

		default:
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

			cfg.Jobs = append(cfg.Jobs, j)
		}
	}

	if len(cfg.Jobs) == 0 {
		return nil, fmt.Errorf("Please add atleast one job to the config")

	}

	if cfg.Mailer.Id == "" && cfg.Slack.Webhook == "" {
		return nil, fmt.Errorf("Please add either Mailer or Slack to the config")
	}

	err = raw.StrictMapTo(cfg)
	if err != nil {
		return nil, fmt.Errorf("Error while reading config: %v", err)
	}

	return cfg, nil
}
