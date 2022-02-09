package controller

import (
	"github.com/sirupsen/logrus"
	"strings"
)

func EvalFunction(in string, errorWords []string) (healthy bool, err error) {
	logrus.Infoln("> EvalFunction")
	healthy = true
	for _, ew := range errorWords {
		if strings.Contains(in, ew) {
			healthy = false
		}
		logrus.WithFields(logrus.Fields{
			"errorWord": ew,
			"healthy":   healthy,
		}).Infoln("next ew")
	}
	logrus.WithField(
		"healthy", healthy).Infoln(">> OK")
	logrus.Infoln("")
	return
}
