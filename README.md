# Pingmon

[![Go Reference](https://pkg.go.dev/badge/github.com/p-tupe/pingmon.svg)](https://pkg.go.dev/github.com/p-tupe/pingmon)

> Under construction

Pingmon is a cli tool for site monitoring and downtime alert service.

## Setup

Pingmon can be installed in several ways - as a docker container, as a systemd unit, as a binary, or even build from source. Here's a direct installation from pkg.go.dev:

```sh
go install github.com/p-tupe/pingmon@latest
```

Give it a `config.json` file, and off it goes

```sh
pingmon -c /path/to/config.json
```

## Configuration

```json
{
  "sites": [
    { "url": "https://www.example.com" },
    { "url": "https://google.com" }
  ],

  "ntfy": "ntfy.sh/channel",

  "slackWebhook": "https://hooks.slack.com/services/channel123ID",

  "mailTo": ["receipient@one.com", "receipient@two.com"],
  "mailer": {
    "host": "smtp.mail.com",
    "port": 587,
    "username": "mailerUsername",
    "password": "mailerPassword",
    "from": "from@mail.com"
  }
}
```

## How it works

Pingmon sends an http request to the site URLs defined in the config at the start of every interval. On receiving a response, it parses the headers and extracts the status code.

If the status code is less than 500 (0 - 499), the site is considered up.

If the site is up, then the status is simply saved in DB. If the site is down however, then an alert is generated on the configured channels, depending on the config values.
