package controller

import (
	"github.com/gregdel/pushover"
	"github.com/just1689/pushover-cmd-daemon/model"
	"github.com/sirupsen/logrus"
)

func NotifyPushoverFromConfig(c *model.Config) error {
	logrus.Infoln("> NotifyPushoverFromConfig")
	app := pushover.New(c.PushoverToken)
	recipient := pushover.NewRecipient(c.PushoverRecipient)
	message := pushover.NewMessageWithTitle(c.MsgBody, c.MsgTitle)
	message.Priority = c.MsgPriority
	if _, err := app.SendMessage(message, recipient); err != nil {
		logrus.Errorln(err)
		return err
	}
	logrus.Infoln(">> OK\n")
	return nil
}
