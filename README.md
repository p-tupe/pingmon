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

## Usage

```sh
pingmon help
pingmon init
pingmon test
pingmon status
pingmon log
pingmon start
pingmon stop
pingmon remove
```

## Config

A sample config:

```ini
site.com=30
email=sender@mail.com
slack_webhook=https://something
up_status=
down_status=
```

## How it works

Pingmon sends an http request to the site URLs defined in the config at the start of every interval. On receiving a response, it parses the headers and extracts the status code.

If the status code is less than 500 (0 - 499), the site is considered UP. This behaviour is configurable. Is the site is UP, then the status is simply logged.

If the site is DOWN, then an alert is generated on either the email or slack medium, depending on the config values.

All config options are global by default, and thus apply to all site. However, each option can also pe configured per site.
