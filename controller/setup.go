package controller

import (
	"github.com/just1689/pushover-cmd-daemon/model"
	"github.com/sirupsen/logrus"
)

func Setup(c *model.Config) {
	logrus.Infoln("> Setting up cronjob", c.Cron)
	InitiateCron(c.Cron, func() {
		logrus.Infoln("> cron function run")
		healthy, err := Check(c)
		if err != nil || !healthy {
			NotifyPushoverFromConfig(c)
		}
		logrus.Infoln(">> OK")
		logrus.Infoln("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~\n\n")
	})
	logrus.Infoln(">> OK\n")

}