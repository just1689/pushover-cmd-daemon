package controller

import (
	"github.com/robfig/cron"
)

func InitiateCron(cronString string, f func()) {
	c := cron.New()
	c.AddFunc(cronString, f)
	c.Start()
}
