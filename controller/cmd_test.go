package controller

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"testing"
)

func TestCMD(t *testing.T) {
	out, err := CMD("echo", []string{"1"})
	if err != nil {
		logrus.Errorln(err)
		t.Fail()
	}
	if len(out) < 1 || out[0] != '1' {
		t.Error("out is empty for CMD(echo 1)")
		t.Error(fmt.Sprintf("'%s'", out))
		t.Fail()
	}
}
