package main

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	i "github.com/p-tupe/pingmon/internal"
)

func TestNewConfig(t *testing.T) {
	cases := []struct {
		file string
		cfg  i.Config
		err  error
	}{
		{"noSite.json", i.Config{}, i.ErrNoSite},
		{"noAlert.json", i.Config{Sites: []i.Site{{URL: "example.com"}}}, i.ErrNoAlert},
		{"noRecipient.json", i.Config{Sites: []i.Site{{URL: "example.com"}, {URL: "google.com"}}, Mailer: &i.Mailer{}}, i.ErrNoMailTo},
		{"onlyPostAlert.json", i.Config{Sites: []i.Site{{URL: "example.com"}}, PostRequest: &i.PostRequest{}}, nil},
		{"mailerAndMailTo.json", i.Config{Sites: []i.Site{{URL: "example.com"}}, Mailer: &i.Mailer{}, MailTo: []string{"mail@to.com"}}, nil},
		{"intervalInSite.json", i.Config{Sites: []i.Site{{Interval: 30, URL: "example.com"}}, PostRequest: &i.PostRequest{}}, nil},
		{"fullPostRequest.json", i.Config{Sites: []i.Site{{URL: "example.com"}}, PostRequest: &i.PostRequest{URL: "request.com", ContentType: "application/json"}}, nil},
	}

	for _, c := range cases {
		t.Run(c.file, func(t *testing.T) {
			configPath := filepath.Join(t.TempDir(), c.file)
			configJson, err := json.Marshal(c.cfg)
			if err != nil {
				panic(err)
			}
			os.WriteFile(configPath, []byte(configJson), 0644)

			parsedConfig, err := i.NewConfig(configPath)
			if errors.Is(err, c.err) {
				return
			}

			if err != nil {
				t.Error(err)
			}

			if !reflect.DeepEqual(parsedConfig, c.cfg) {
				t.Errorf("Expected %v, Got %v", c.cfg, parsedConfig)
			}
		})
	}
}
