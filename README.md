# Pingmon

> Under construction

A simple fire 'n forget site monitoring service.

## Setup

```sh
go install github.com/EMPAT94/pingmon@latest
```

Pingmon works off a config file and sets itself up as a systemd service. Just `pingmon setup` and follow the prompts to get up and running.

## Help

```
Usage:
    pingmon [command]

Commands:
    help       Prints this help
    setup      Setup pingmon config and service files
    status     Show latest pingmon results
    remove     Clear all pingmon data and stop service
    test       Test network and notifications
    log        Show latest pingmon logs
    start      Start pingmon service, if not already running
    stop       Clear all pingmon data and stop service
```

## How it works

Pingmon sends an http request to the site URLs defined in the config at the start of every interval. On receiving a response, it parses the headers and extracts the status code.

If the status code is less than 500 (0 - 499), the site is considered UP. This behaviour is configurable per site.

If the site is UP, then the status is simply logged. If the site is DOWN however, then an alert is generated on either the email or slack medium, depending on the config values.

All config options are global by default, and thus apply to all site. However, each option can also be configured per site.
