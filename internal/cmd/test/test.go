package test

import (
	"fmt"

	"github.com/p-tupe/pingmon/internal/app/config"
)

func Test(config *config.Config) {
	fmt.Printf("%+v", config)
	// email.SendMail(config)
}
