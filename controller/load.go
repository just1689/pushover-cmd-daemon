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

// LoadFromFile attempts to read and unmarshal the config
func LoadFromFile(f string) (*model.Config, error) {
	c := &model.Config{}
	b, err := ioutil.ReadFile(f)
	if err != nil {
		logrus.Errorln("could not load file:", f)
		return c, errors.New(model.ErrorNoConfig)
	}
	if err = json.Unmarshal(b, c); err != nil {
		logrus.Errorln("could not unmarshal file to config:", f)
		return c, err
	}
	return c, err
}

// CreateConfig will generate the config.json and write the file
func CreateConfig() {
	b, err := json.Marshal(model.GetConfigDefaults())
	if err != nil {
		logrus.Errorln(err)
		return
	}
	ioutil.WriteFile("config.json", b, 0645)
	return
}
