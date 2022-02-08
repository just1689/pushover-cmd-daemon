package controller

import (
	"testing"
)

var TestPushOver = false

func TestNotifyPushoverFromConfig(t *testing.T) {
	if !TestPushOver {
		return
	}
	c, err := LoadFromFile("../config.json")
	if err != nil {
		t.Error("could not load config.json")
		t.Fail()
	}
	if err = NotifyPushoverFromConfig(c); err != nil {
		t.Fail()
	}
}
