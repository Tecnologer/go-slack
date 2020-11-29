# Slack API Wrapper

Slack API wrapper with Go

## Example

```golang
    var s *slack.Slack

    func main() {
        s = slack.New("<slackToken>")

        err := s.SetCommand("hola", sayHi)
        if err != nil {
            log.Error("set command hola: %v")
        }

        s.StartWithWebhook(8080)
    }

    func sayHi(event *response.EventResponse) {
        err := s.SendTextMessage(event.GetChannel(), "hola desde go")
        if err != nil {
            logrus.WithError(err).Error("sending message")
        }
    }
```
