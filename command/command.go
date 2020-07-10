package command

import (
	"strconv"
	"strings"
)

// Response a response wrapper
type Response func(string)

// MessageEvent sends the relevant params to the command handling the message event.
type MessageEvent struct {
	MessageID    string
	Channel      string
	Username     string
	Message      string
	IsModerator  bool
	IsSubscriber bool
	Tags         map[string]string
}

// Command basic command interface
type Command interface {
	Execute(Response, MessageEvent)
}

// Config basic command config struct
type Config struct {
	Enabled bool   `json:"enabled"`
	ID      string `json:"id"`
	Key     string `json:"key"`
}

// Format replace $ tokens in output
func (m MessageEvent) Format(s string) string {
	rTokens := strings.Split(s, " ")
	mTokens := strings.Split(m.Message, " ")
	output := []string{}

	for _, t := range rTokens {
		key := t[1:]
		i, err := strconv.Atoi(key)
		val, ok := m.Tags[key]

		switch {
		case t[0] != '$':
			output = append(output, t)
		case err == nil:
			output = append(output, mTokens[i-1])
		case ok:
			output = append(output, val)
		default:
			output = append(output, t)
		}
	}

	return strings.Join(output, " ")
}
