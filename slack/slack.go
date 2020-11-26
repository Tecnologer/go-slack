package slack

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/tecnologer/go-slack/slack/models/response"
)

const apiURL = "https://slack.com/api/%s"

//Slack is the base struct for slack data
type Slack struct {
	CmdPrefix string
	token     string
	allMsg    func(*response.Message)
	commands  map[string]func(*response.Message)
}

//New creates new instance of Slack
func New(token string) *Slack {
	return &Slack{
		CmdPrefix: "/",
		token:     token,
		allMsg:    nil,
		commands:  make(map[string]func(*response.Message)),
	}
}

func (s *Slack) getEndPoint(action string) string {
	return fmt.Sprintf(apiURL, action)
}

func (s *Slack) SendTextMessage(channel, msg string) error {
	endpoint := s.getEndPoint("chat.postMessage")

	v := url.Values{}

	v.Add("token", s.token)
	v.Add("text", msg)
	v.Add("channel", channel)

	res, err := http.PostForm(endpoint, v)

	if err != nil {
		return errors.Wrap(err, "send message post")
	}

	body := &response.SendMessage{}
	bodyDecoder := json.NewDecoder(res.Body)
	err = bodyDecoder.Decode(body)
	if err != nil {
		return errors.Wrap(err, "decode send message body: parsing response (json)")
	}

	if !body.OK {
		var genericBody interface{}
		_ = bodyDecoder.Decode(genericBody)
		return fmt.Errorf("slack error response: %v", genericBody)
	}

	return nil
}

func (s *Slack) StartWithWebhook(url string, port int) {
	chanUpdates, err := s.setWeebhook(url, port)
	if err != nil {
		logrus.WithError(err).Error("register for updates")
		return
	}
	for update := range chanUpdates {
		s.validateCommand(update)

		if s.allMsg != nil {
			s.allMsg(update)
		}
	}
}
