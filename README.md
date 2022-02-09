# pushover-cmd-daemon

![Main status](https://github.com/just1689/pushover-cmd-daemon/actions/workflows/go.yml/badge.svg)
<a href="https://goreportcard.com/report/github.com/just1689/pushover-cmd-daemon"><img src="https://goreportcard.com/badge/github.com/just1689/pushover-cmd-daemon"></a>
<a href="https://codebeat.co/projects/github-com-just1689-pushover-cmd-daemon-main"><img alt="codebeat badge" src="https://codebeat.co/badges/8dcd25bc-51d4-4624-9697-aa17fc58ce3b" /></a>
<a href="https://codeclimate.com/github/just1689/pushover-cmd-daemon/maintainability"><img src="https://api.codeclimate.com/v1/badges/0463e4ecba35fc60b3d7/maintainability" /></a>

Simple background app that evaluates the result of executing a command &amp; alerts via Pushover

## Setup

This project uses github.com/robfig/cron for the cron. See cron docs
here: https://pkg.go.dev/github.com/robfig/cron?utm_source=godoc

- To generate the config file just run the executable for the first time. It will generate a file `config.json`.
- Run the binary in the same directory as `config.json`.

## Config

The `config.json` contains the configuration for the app.

```json
{
  "cron": "@every 1h",
  "cmd": "print",
  "args": [
    "1"
  ],
  "errorWords": [
    "error"
  ],
  "msgTitle": "Error",
  "msgBody": "Could not find 1 in result",
  "msgPriority": 0,
  "pushoverToken": "<your token here>",
  "pushoverRecipient": "<your recipient here>"
}
```

## To do

- [x] Config model
- [x] Run CMD
- [x] Setup cron
- [x] Load result from CMD run
- [x] Run error function
- [x] If error func true, alert
- [x] Call Pushover notification
- [ ] Manage releases with Github Actions
- [ ] Release binary for Linux, MacOS and Windows
- [ ] Better logging strategy
- [ ] Param for config file name.
- [ ] Param for ANSI color logs
- [ ] Timeouts, threading.
- [ ] Watch for hanging and notify.
