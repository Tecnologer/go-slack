package slack

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/sirupsen/logrus"
	"github.com/tecnologer/go-slack/slack/models/response"
)

var msgChannel chan *response.EventResponse

func (s *Slack) setWeebhook(port int) (chan *response.EventResponse, error) {
	msgChannel = make(chan *response.EventResponse)
	go http.ListenAndServe(fmt.Sprintf(":%d", port), http.HandlerFunc(webhookReceiver))
	return msgChannel, nil
}

func webhookReceiver(res http.ResponseWriter, req *http.Request) {
	res.Header().Add("Content-Type", "application/json")
	enableEvent := &response.EnableEventRequest{}
	event := &response.EventResponse{}
	reqBody, err := ioutil.ReadAll(req.Body)

	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte("invalid request"))
		return
	}

	err = json.Unmarshal(reqBody, enableEvent)

	if err != nil {
		logrus.WithError(err).Warning("parse body to enable event")
	}

	if err == nil && enableEvent.Challenge != "" {
		res.Write([]byte(enableEvent.Challenge))
		return
	}
	logrus.Debug(string(reqBody))
	err = json.Unmarshal(reqBody, event)
	if err != nil {
		logrus.WithError(err).Error("parse body to event")
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	if event != nil {
		logrus.Debugf("New Message: %s", event.GetMessageTxt())
		msgChannel <- event
	}
}
