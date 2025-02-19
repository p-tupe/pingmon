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

func checkCustomPath(customPath string) error {
	_, err := os.Stat(customPath)
	if err != nil {
		return err
	}
	return nil
}

func createNewConfigFile() (string, error) {
	configFilePath, err := xdg.ConfigFile("pingmon/config.ini")
	if err != nil {
		return "", err
	}
	_, err = os.Create(configFilePath)
	if err != nil {
		return "", err
	}
	return configFilePath, nil
}

func ReadConfig(customPath string) (*Config, error) {
	cfg := new(Config)

	configFilePath := ""

	if customPath != "" {
		err := checkCustomPath(customPath)
		if err != nil {
			return nil, fmt.Errorf("Error while reading custom config file path: %v", err)
		}
		configFilePath = customPath
	} else {
		var err error
		configFilePath, err = xdg.SearchConfigFile("pingmon/config.ini")
		if err != nil {
			if strings.HasPrefix("could not locate", err.Error()) {
				configFilePath, err = createNewConfigFile()
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
