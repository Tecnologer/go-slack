package slack

import (
	"fmt"
	"strings"

	"github.com/tecnologer/go-slack/slack/models/response"
)

func (s *Slack) SetCommand(cmd string, action func(*response.EventResponse)) error {
	if action == nil {
		return fmt.Errorf("callback function is required for a command")
	}

	if s.commands == nil {
		s.commands = make(map[string]func(*response.EventResponse))
	}

	if !strings.HasPrefix(cmd, s.CmdPrefix) {
		cmd = s.CmdPrefix + cmd
	}
	s.commands[cmd] = action

	return nil
}

func (s *Slack) validateCommand(update *response.EventResponse) {
	cmd := getCmdFromMsg(update)
	if cmd == "" {
		return
	}

	if action, exists := s.commands[cmd]; exists {
		action(update)
	}
}

func getCmdFromMsg(msg *response.EventResponse) string {
	if msg == nil {
		return ""
	}

	msgParts := strings.Split(strings.Trim(msg.GetMessageTxt(), " "), " ")
	if len(msgParts) == 0 {
		return ""
	}

	return msgParts[0]
}
