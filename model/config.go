package model

type Config struct {
	Cron              string   `json:"cron"`
	CMD               string   `json:"cmd"`
	Args              []string `json:"args"`
	ErrorWords        []string `json:"errorWords"`
	MsgTitle          string   `json:"msgTitle"`
	MsgBody           string   `json:"msgBody"`
	MsgPriority       int      `json:"msgPriority"`
	PushoverToken     string   `json:"pushoverToken"`
	PushoverRecipient string   `json:"pushoverRecipient"`
}
