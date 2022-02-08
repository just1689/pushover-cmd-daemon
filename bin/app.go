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

	controller.Setup(loadConfig())

	logrus.Infoln("> Listening for SIGTERM\n\n")
	sigChan := make(chan os.Signal)
	fmt.Println("\nShutting down\n", <-sigChan)

}

func loadConfig() (c *model.Config) {
	var err error
	c, err = controller.Load()
	if err != nil {
		if err.Error() == model.ErrorNoConfig {
			logrus.Infoln(">> No config found. Generating")
			controller.CreateConfig()
			os.Exit(0)
		}
		os.Exit(1)
		return
	}
	logrus.WithField("err", err).Infoln("> Load()")
	return c
}

func setupLogger() {
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
}
