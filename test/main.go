package main

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/tecnologer/go-secrets"
	"github.com/tecnologer/go-secrets/config"
	"github.com/tecnologer/go-slack/slack"
	"github.com/tecnologer/go-slack/slack/models/response"
)

var s *slack.Slack

func main() {
	secrets.InitWithConfig(&config.Config{})
	slackToken := secrets.GetKeyString("slack.token")

	s = slack.New(slackToken)

	err := s.SetCommand("hola", sayHi)
	if err != nil {
		logrus.WithError(err).Error("set command hola")
	}

	fmt.Println(s)
}

func sayHi(msg *response.Message) {
	err := s.SendTextMessage("ULC4RM8HG", "hola desde go")
	if err != nil {
		logrus.WithError(err).Error("sending message")
	}
}
