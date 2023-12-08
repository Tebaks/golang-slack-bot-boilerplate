package sayhi

import (
	"app/pkg/handler/slack/slashcommand"
	"fmt"
	"github.com/slack-go/slack"
	"github.com/slack-go/slack/socketmode"
)

type handler struct {
}

func NewHandler() slashcommand.Handler {
	return handler{}
}

func (h handler) HandleCommand(evt *socketmode.Event, client *socketmode.Client) error {
	data := evt.Data.(slack.SlashCommand)
	_, _, err := client.PostMessage(data.ChannelID, slack.MsgOptionText(fmt.Sprintf("Hi, I'm Test Bot I got your command."), false))
	if err != nil {
		return err
	}

	client.Ack(*evt.Request)
	return nil
}
