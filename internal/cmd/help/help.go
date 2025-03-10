package help

import (
	"fmt"

	"github.com/EMPAT94/pingmon/internal/app/config"
)

func Show() {
	fmt.Print(`
A simple fire 'n forget site monitoring service.

Usage:
    pingmon [command]

Commands: 
`)

	for k, v := range config.Cmds {
		fmt.Printf("    %-10s %s\n", k, v)
	}

	fmt.Println()
}
