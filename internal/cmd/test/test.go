package test

import (
	"github.com/EMPAT94/pingmon/internal/app/email"
	"github.com/EMPAT94/pingmon/internal/config"
)

func Test(config *config.Config) {
	email.SendMail(config)
}
