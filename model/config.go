package model

// Config describes the options for running the system
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

// GetConfigDefaults is a hardcoded set of defaults for generating the config.json
func GetConfigDefaults() Config {
	return Config{
		Cron:              "@every 1m",
		CMD:               "print",
		Args:              []string{"1"},
		ErrorWords:        []string{"error"},
		MsgTitle:          "Error",
		MsgBody:           "Could not find 1 in result",
		MsgPriority:       0,
		PushoverToken:     "<your token>",
		PushoverRecipient: "<recipient>",
	}
}
