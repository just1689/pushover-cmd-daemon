# pushover-cmd-daemon

Simple background app that evaluates the result of executing a command &amp; alerts via Pushover


## Setup

This project uses github.com/robfig/cron for the cron. See cron docs here: https://pkg.go.dev/github.com/robfig/cron?utm_source=godoc


- To generate the config file just run the executable for the first time. It will generate a file `config.json`.
- Run the binary in the same directory as `config.json`.




## To do
- [x] Config model
- [x] Run CMD
- [x] Setup cron
- [x] Load result from CMD run
- [x] Run error function
- [x] If error func true, alert
- [x] Call Pushover notification
- [ ] Timeouts, threading.
- [ ] Watch for hanging and notify.
