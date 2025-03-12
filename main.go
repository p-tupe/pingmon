package main

import (
	"flag"
	"log"

	"github.com/EMPAT94/pingmon/internal/app/config"

	"github.com/EMPAT94/pingmon/internal/cmd/help"
	"github.com/EMPAT94/pingmon/internal/cmd/setup"
	"github.com/EMPAT94/pingmon/internal/cmd/start"
	"github.com/EMPAT94/pingmon/internal/cmd/test"
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
