package controller

import (
	"github.com/just1689/pushover-cmd-daemon/model"
	"github.com/sirupsen/logrus"
)

func Check(c *model.Config) (healthy bool, err error) {
	var in string
	in, err = CMD(c.CMD, c.Args)
	healthy, err = EvalFunction(in, c.ErrorWords)
	logrus.WithField("healthy", healthy).Infoln("> Starting Check")
	return
}
