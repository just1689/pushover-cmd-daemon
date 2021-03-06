package controller

import (
	"github.com/sirupsen/logrus"
	"os/exec"
)

// CMD runs a linux command and returns the full output
func CMD(cmd string, args []string) (out string, err error) {
	logrus.Infoln("> Starting CMD")
	var b []byte
	b, err = exec.Command(cmd, args...).Output()
	if err != nil {
		logrus.Errorln(err)
		logrus.Errorln(string(b))
		return
	}
	out = string(b)
	logrus.Infoln(">> OK")
	logrus.Infoln("")
	return
}
