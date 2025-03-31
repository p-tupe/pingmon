/*
Pingmon is a cli tool for site monitoring and downtime alert service.

It pulls mailer/slack info and the sites list from a config file,
and checks the status at every interval (default 30min).
If the response status code is >499 or there was an error reach the site,
it sends an alert in the form of an email or slack message as set in the config.

Usage:

	pingmon [command]

Commands:

	help       Prints this help
	setup      Setup pingmon config and service files
	status     Show latest pingmon results
	remove     Clear all pingmon data and stop service
	test       Test network and notifications
	log        Show latest pingmon logs
	start      Start pingmon service, if not already running
	stop       Clear all pingmon data and stop service

Pingmon can be installed in several ways - as a docker container, as a systemd unit, as a binary, or build from source.

For convenience, a "setup" command is provided that creates a user-guided config.json file in a suitable directory depending on your system, and if it detects a linux OS, will create a corresponding systemd service file as well.
*/
package main

import (
	"flag"
	"log"

	"github.com/p-tupe/pingmon/internal/app/config"
	"github.com/p-tupe/pingmon/internal/cmd/help"
	"github.com/p-tupe/pingmon/internal/cmd/setup"
	"github.com/p-tupe/pingmon/internal/cmd/start"
	"github.com/p-tupe/pingmon/internal/cmd/test"
)

func main() {
	flag.Parse()
	var cmd = flag.Arg(0)

	// TODO: Add SIG handling

	switch cmd {
	case "help":
		help.Show()
	case "setup":
		setup.Setup()
	case "start":
		start.Start()
	case "test":
		config, err := config.Parse()
		if err != nil {
			log.Fatalln("Error reading config.json", err)
		}
		test.Test(config)
	default:
		help.Show()
	}
}
