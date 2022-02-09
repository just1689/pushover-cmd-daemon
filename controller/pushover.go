package controller

import (
	"github.com/gregdel/pushover"
	"github.com/just1689/pushover-cmd-daemon/model"
	"github.com/sirupsen/logrus"
)

var (
	// PushoverClient communicates with the Pushover API
	PushoverClient *pushover.Pushover
)

func NewPushOverClient(token string) *pushover.Pushover {
	logrus.Infoln("> NewPushOverClient")
	return pushover.New(token)
}

func GeneratePushover(recipientID string, title, body string, priority int) (*pushover.Recipient, *pushover.Message) {
	logrus.Infoln("> GeneratePushover")
	recipient := pushover.NewRecipient(recipientID)
	message := pushover.NewMessageWithTitle(body, title)
	message.Priority = priority
	return recipient, message
}

func NotifyPushoverFromConfig(c *model.Config) error {
	logrus.Infoln("> NotifyPushoverFromConfig")
	recipient, message := GeneratePushover(c.PushoverRecipient, c.MsgTitle, c.MsgBody, c.MsgPriority)
	return NotifyPushover(recipient, message)
}
func NotifyPushoverFromSystemError(c *model.Config, errStr string) error {
	logrus.Infoln("> NotifyPushoverFromSystemError")
	recipient, message := GeneratePushover(c.PushoverRecipient, "Faulty - "+c.MsgTitle, errStr, c.MsgPriority)
	return NotifyPushover(recipient, message)
}

func NotifyPushover(recipient *pushover.Recipient, message *pushover.Message) error {
	logrus.Infoln("> NotifyPushover")
	if _, err := PushoverClient.SendMessage(message, recipient); err != nil {
		logrus.Errorln(err)
		return err
	}
	logrus.Infoln(">> OK")
	logrus.Infoln("")
	return nil
}
