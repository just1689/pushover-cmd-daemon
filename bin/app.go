package main

import (
	"fmt"
	"github.com/just1689/pushover-cmd-daemon/controller"
	"github.com/just1689/pushover-cmd-daemon/model"
	"github.com/sirupsen/logrus"
	"os"
)

func main() {
	setupLogger()
	logrus.Infoln("")
	logrus.Infoln("~~~ Starting Pushover CMD Daemon ~~~")
	logrus.Infoln("")

	c := loadConfig()
	controller.PushoverClient = controller.NewPushOverClient(c.PushoverToken)
	controller.Setup(c)

	logrus.Infoln("> Listening for SIGTERM")
	logrus.Infoln("")
	sigChan := make(chan os.Signal)
	logrus.Infoln("")
	fmt.Println("Shutting down", <-sigChan)
	logrus.Infoln("")

}

func loadConfig() (c *model.Config) {
	var err error
	c, err = controller.Load()
	if err.Error() == model.ErrorNoConfig {
		logrus.Infoln(">> No config found. Generating")
		controller.CreateConfig()
		os.Exit(0)
	} else if err != nil {
		logrus.Fatalln(err)
	}
	logrus.WithField("err", err).Infoln("> Load()")
	return c
}

func setupLogger() {
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
}
