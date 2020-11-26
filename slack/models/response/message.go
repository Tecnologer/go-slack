package response

type SendMessage struct {
	OK      bool     `json:"ok"`
	Channel string   `json:"channel"`
	Message *Message `json:"message"`
}

type Message struct {
	BotID string `json:"bot_id"`
	Type  string `json:"message"`
	Text  string `json:"text"`
}

/*
{
    "ok": true,
    "channel": "D01F52GPGBH",
    "ts": "1606365112.000400",
    "message": {
        "bot_id": "B01FKQ9BRDY",
        "type": "message",
        "text": "hola",
        "user": "U01FD2SE094",
        "ts": "1606365112.000400",
        "team": "T045T6RV5",
        "bot_profile": {
            "id": "B01FKQ9BRDY",
            "deleted": false,
            "name": "tecnologer-bot",
            "updated": 1606364282,
            "app_id": "A01FRML3SMA",
            "icons": {
                "image_36": "https:\/\/a.slack-edge.com\/80588\/img\/plugins\/app\/bot_36.png",
                "image_48": "https:\/\/a.slack-edge.com\/80588\/img\/plugins\/app\/bot_48.png",
                "image_72": "https:\/\/a.slack-edge.com\/80588\/img\/plugins\/app\/service_72.png"
            },
            "team_id": "T045T6RV5"
        }
    }
}
*/
