package controller

import (
	"github.com/robfig/cron"
)

// InitiateCron starts the cron that runs the function
func InitiateCron(cronString string, f func()) {
	c := cron.New()
	c.AddFunc(cronString, f)
	c.Start()
}
