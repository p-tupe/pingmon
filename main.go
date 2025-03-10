package main

import (
	"flag"

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
		setup.Init()
	case "start":
		start.Start()
	case "test":
		var config = config.Parse()
		test.Test(config)
	default:
		help.Show()
	}
}
