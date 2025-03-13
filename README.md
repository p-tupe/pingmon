# Pingmon

[![Go Reference](https://pkg.go.dev/badge/github.com/EMPAT94/pingmon.svg)](https://pkg.go.dev/github.com/EMPAT94/pingmon)

> Under construction

Pingmon is a cli tool for site monitoring and downtime alert service.

## Setup

Pingmon can be installed in several ways - as a docker container, as a systemd unit, as a binary, or even build from source:

```sh
go install github.com/EMPAT94/pingmon@latest
```

```sh
curl https://.../install.sh | sh
```

```sh
docker pull ...
```

For convenience, a "setup" command is provided that creates a user-guided config.json file in a suitable directory depending on your system, and if it detects a linux OS, will create a corresponding systemd service file as well.

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

## Advanced Config

## How pingmon works

Pingmon sends an http request to the site URLs defined in the config at the start of every interval. On receiving a response, it parses the headers and extracts the status code.

If the status code is less than 500 (0 - 499), the site is considered UP. This behaviour is configurable per site.

If the site is UP, then the status is simply logged. If the site is DOWN however, then an alert is generated on either the email or slack medium, depending on the config values.

All config options are global by default, and thus apply to all site. However, each option can also be configured per site.
