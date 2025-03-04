package app

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

var l log.Logger = *log.Default()

func Ping(site string) error {
	l.Printf("[INFO] Pinging %s", site)

	res, err := http.Get(site)
	if err != nil {
		return err
	}

	l.Printf("[INFO] Status for %s %v", site, res.StatusCode)

	ok := (res.StatusCode > 199 && res.StatusCode < 299) ||
		(res.StatusCode >= 400 && res.StatusCode < 499)

	if !ok {
		return fmt.Errorf("Site returned a status of %s", res.Status)
	}

	return nil
}

func Alert(site string, statusCode int, url string) {
	msg := fmt.Sprintf(`{"text":"Could no reach %s, received status code %d. It may be down!"}`, site, statusCode)
	_, err := http.Post(url, "application/json", strings.NewReader(msg))
	if err != nil {
		l.Print("[ERROR] " + err.Error())
	}
}
