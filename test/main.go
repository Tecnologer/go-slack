package main

import (
	"flag"
	"fmt"
	"regexp"
	"time"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/tecnologer/go-secrets"
	"github.com/tecnologer/go-secrets/config"
	"github.com/tecnologer/go-slack/slack"
	"github.com/tecnologer/go-slack/slack/models/response"
)

var debug = flag.Bool("debug", false, "set debug level for logs")
var port = flag.Int("port", 8088, "port webhook")
var s *slack.Slack
var customAction map[*regexp.Regexp]func(*response.EventResponse)

func main() {
	flag.Parse()
	if *debug {
		logrus.SetLevel(logrus.DebugLevel)
	}
	customAction = make(map[*regexp.Regexp]func(*response.EventResponse))

	secrets.InitWithConfig(&config.Config{})
	slackToken := secrets.GetKeyString("slack.token")

	s = slack.New(slackToken)

	err := s.SetCommand("hola", sayHi)
	if err != nil {
		logrus.WithError(err).Error("set command hola")
	}
	_ = s.SetCommand("hora", getTime)
	configCustomAction()

	s.AllEvents(allEvents)
	s.StartWithWebhook(*port)
}

func configCustomAction() {
	actionsMap := map[string]func(*response.EventResponse){
		`^hola.*`:                             sayHi,
		`(?mi)(dime la hora|¿?que hora es\?)`: getTime,
	}

	for expr, action := range actionsMap {
		err := addCustomAction(expr, action)
		if err != nil {
			logrus.WithError(err).WithField("expr", expr).Error("set custom action")
		}
	}
}

func sayHi(event *response.EventResponse) {
	err := s.SendTextMessage(event.GetChannel(), "hola desde go")
	if err != nil {
		logrus.WithError(err).Error("sending message")
	}
}

func getTime(event *response.EventResponse) {
	args := event.GetCmdArgs("/hora")
	isUtc := len(args) > 0 && args[0] == "utc"
	t := time.Now()
	if isUtc {
		t = time.Now().UTC()
	}
	ts := t.Format("Mon, 02 Jan 2006 03:04:05.999 PM -07:00")
	err := s.SendTextMessage(event.GetChannel(), fmt.Sprintf("La hora es: %s", ts))
	if err != nil {
		logrus.WithError(err).Error("sending message")
	}
}

func addCustomAction(expr string, action func(*response.EventResponse)) error {
	reg, err := regexp.Compile(expr)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("error creating custom action for expresion: %s", expr))
	}

	customAction[reg] = action

	return nil
}

func allEvents(event *response.EventResponse) {
	for regex, action := range customAction {
		if regex.Match([]byte(event.GetMessageTxt())) {
			action(event)
		}
	}
}
