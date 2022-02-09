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
	logrus.Infoln("\n\n~~~ Starting Pushover CMD Daemon ~~~\n\n")

	c := loadConfig()
	controller.PushoverClient = controller.NewPushOverClient(c.PushoverToken)
	controller.Setup(c)

	logrus.Infoln("> Listening for SIGTERM\n\n")
	sigChan := make(chan os.Signal)
	fmt.Println("\nShutting down\n", <-sigChan)

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
