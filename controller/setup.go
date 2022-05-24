package controller

import (
	"github.com/just1689/pushover-cmd-daemon/model"
	"github.com/sirupsen/logrus"
)

var CurrentHealthStatus = true

func Setup(c *model.Config) {
	logrus.Infoln("> Setting up cronjob", c.Cron)
	InitiateCron(c.Cron, func() {
		logrus.Infoln("> cron function run")
		healthy, err := Check(c)
		if err != nil {
			NotifyPushoverFromSystemError(c, err.Error())
			return
		}
		if !healthy {
			NotifyPushoverFromConfig(c)
			CurrentHealthStatus = false
			return
		}
		CurrentHealthStatus = true
		logrus.Infoln(">> OK")
		logrus.Infoln("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
		logrus.Infoln("")
	})
	logrus.Infoln(">> OK")
	logrus.Infoln("")

}
