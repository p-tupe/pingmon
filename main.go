package main

import (
	"fmt"
	"log"

	"github.com/EMPAT94/pingmon/internal/config"
)

func main() {
	l := log.Default()

	c, err := config.ReadConfig("~/.config/pingmon/config.ini")
	if err != nil {
		l.Fatalln(err)
	}

	fmt.Println("Main:", c)
}
