# Pingmon

> Under construction

A simple fire 'n forget site monitoring service.

## Setup

```sh
go install .../pingmon.git
```

OR

```sh
curl .../install.sh
```

Pingmon works off a config file and sets itself up as a service. Following commands, in order, help setting things up the first time:

1. `pingmon init` - Create config and systemd files, if they don't already exits. This will ask for some inputs to populate the config.

2. `pingmon test` - A one off version of "start", will send test notifications and show site response as output.

3. `pingmon start` - Enable monitoring service, if everything worked fine above.

## Help

```
Usage:
    pingmon [command]

Commands:
    help       Prints this help
    init       Initialize pingmon site monitoring service
    status     Show latest pingmon results
    remove     Clear all pingmon data and stop service
    test       Test network and notifications
    log        Show latest pingmon logs
    start      Start pingmon service, if not already running
    stop       Clear all pingmon data and stop service
```

## Config

A sample config:

```ini
site.com=30
slack_webhook=https://something
host="smtp.something.com"
port="25/587/465"
username="username"
password="password"
from="sender@mail.com"
to="receipient@mail.com"

```

## How it works

Pingmon sends an http request to the site URLs defined in the config at the start of every interval. On receiving a response, it parses the headers and extracts the status code.

If the status code is less than 500 (0 - 499), the site is considered UP. This behaviour is configurable. Is the site is UP, then the status is simply logged.

If the site is DOWN, then an alert is generated on either the email or slack medium, depending on the config values.

All config options are global by default, and thus apply to all site. However, each option can also pe configured per site.
