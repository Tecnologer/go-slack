package response

import "strings"

type Event struct {
	BotID       string   `json:"bot_id"`
	ClientMsgID string   `json:"client_msg_id"`
	Type        string   `json:"type"`
	Text        string   `json:"text"`
	User        string   `json:"user"`
	Ts          string   `json:"ts"`
	Team        string   `json:"team"`
	Blocks      []*Block `json:"blocks"`
	Channel     string   `json:"channel"`
	EventTs     string   `json:"event_ts"`
	ChannelType string   `json:"channel_type"`
}

//EventResponse strutc of message in event
type EventResponse struct {
	Token              string           `json:"token"`
	TeamID             string           `json:"team_id"`
	APIAppID           string           `json:"api_app_id"`
	Event              *Event           `json:"event"`
	Type               string           `json:"type"`
	EventID            string           `json:"event_id"`
	EventTime          int              `json:"event_time"`
	Authorizations     []*Authorization `json:"authorizations"`
	IsExtSharedChannel bool             `json:"is_ext_shared_channel"`
	EventContext       string           `json:"event_context"`
}

type Block struct {
	Type     string          `json:"type"`
	BlockID  string          `json:"block_id"`
	Elements []*BlockElement `json:"elements"`
}

type BlockElement struct {
	Type     string     `json:"type"`
	Elements []*Element `json:"elements"`
}

type Element struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

type Authorization struct {
	EnterpriseID        interface{} `json:"enterprise_id"`
	TeamID              string      `json:"team_id"`
	UserID              string      `json:"user_id"`
	IsBot               bool        `json:"is_bot"`
	IsEnterpriseInstall bool        `json:"is_enterprise_install"`
}

func (e *EventResponse) GetChannel() string {
	if e.Event == nil {
		return ""
	}
	return e.Event.Channel
}

func (e *EventResponse) GetMessageTxt() string {
	if e.Event == nil {
		return ""
	}
	return e.Event.Text
}

func (e *EventResponse) IsBot() bool {
	if e.Event == nil {
		return false
	}
	return e.Event.BotID != ""
}

func (s *EventResponse) GetCmdArgs(cmd string) []string {
	text := s.GetMessageTxt()

	argsStr := strings.Trim(strings.Replace(text, cmd, "", 1), " ")

	return strings.Split(argsStr, " ")
}
