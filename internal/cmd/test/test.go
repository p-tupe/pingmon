package test

import (
	"github.com/EMPAT94/pingmon/internal/app/config"
	"github.com/EMPAT94/pingmon/internal/app/email"
)

func Test(config *config.Config) {
	email.SendMail(config)
}
