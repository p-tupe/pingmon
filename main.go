package main

import (
	"flag"

	"github.com/EMPAT94/pingmon/internal/cmd/help"
	"github.com/EMPAT94/pingmon/internal/cmd/start"
	"github.com/EMPAT94/pingmon/internal/cmd/test"
	"github.com/EMPAT94/pingmon/internal/config"
)

func main() {
	flag.Parse()
	var cmd = flag.Arg(0)

	switch cmd {
	case "help":
		help.Show()
	case "start":
		start.Start()
	case "test":
		var config = config.Parse()
		test.Test(config)
	default:
		help.Show()
	}
}
