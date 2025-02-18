package main

import (
	"fmt"
	"log"

	"github.com/EMPAT94/pingmon/internal/config"
)

func main() {
	l := log.Default()

	c, err := config.ReadConfig()
	if err != nil {
		l.Fatalln(err)
	}

	fmt.Println("Main:", c)
}
