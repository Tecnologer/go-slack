package slack

//Slack is the base struct for slack data
type Slack struct {
	token string
}

//New creates new instance of Slack
func New(token string) *Slack {
	return &Slack{
		token: token,
	}
}
