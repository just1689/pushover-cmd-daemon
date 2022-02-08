package controller

import (
	"encoding/json"
	"errors"
	"github.com/just1689/pushover-cmd-daemon/model"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
)

func getConfigName() string {
	result := os.Getenv("config")
	if result == "" {
		result = "config.json"
	}
	return result
}

func Load() (*model.Config, error) {
	return LoadFromFile(getConfigName())
}

func LoadFromFile(f string) (*model.Config, error) {
	c := &model.Config{}
	b, err := ioutil.ReadFile(f)
	if err != nil {
		logrus.Errorln("could not load file:", f)
		return c, errors.New(model.ErrorNoConfig)
	}
	err = json.Unmarshal(b, c)
	if err != nil {
		logrus.Errorln("could not unmarshal file to config:", f)
		return c, err
	}
	return c, err
}

func CreateConfig() {
	c := model.Config{
		Cron:              "@every 1m",
		CMD:               "print",
		Args:              []string{"1"},
		ErrorWords:        []string{"error"},
		MsgTitle:          "Error",
		MsgBody:           "Could not find 1 in result",
		MsgPriority:       0,
		PushoverToken:     "xxx",
		PushoverRecipient: "yyy",
	}
	b, err := json.Marshal(c)
	if err != nil {
		logrus.Errorln(err)
		return
	}
	ioutil.WriteFile("config.json", b, 0645)
	return
}
