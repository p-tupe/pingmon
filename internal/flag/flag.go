package flag

import (
	"flag"
	"fmt"
)

var SubCmds = map[string]string{
	"help":    "Prints this help",
	"init":    "Initialize pingmon site monitoring service",
	"test":    "Test network and notifications",
	"status":  "Show latest pingmon results",
	"log":     "Show latest pingmon logs",
	"start":   "Start pingmon service, if not already running",
	"restart": "Gracefully stop pingmon service",
	"stop":    "Clear all pingmon data and stop service",
	"remove":  "Clear all pingmon data and stop service",
}

func Parse() (string, error) {
	flag.Parse()

	x := flag.Arg(0)

	if _, ok := SubCmds[x]; !ok {
		return "", fmt.Errorf("Invalid command, please enter 'pingmon help' for usage.")
	} else {
		return x, nil
	}
}
