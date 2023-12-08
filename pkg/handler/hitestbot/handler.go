package hitestbot

import (
	"app/pkg/configs"
	"app/pkg/handler"
	"github.com/slack-go/slack"
	"github.com/slack-go/slack/slackevents"
	"github.com/slack-go/slack/socketmode"
)

type messageHandler struct {
}

func NewMessageHandler() handler.MessageHandler {
	return messageHandler{}
}

func (h messageHandler) HandleMessage(messageEvent *slackevents.MessageEvent, client *socketmode.Client) error {
	_, _, err := client.PostMessage(messageEvent.Channel, slack.MsgOptionText("Hi, I'm Test Bot I got your message.", false))
	if err != nil {
		return err
	}

	return nil
}

func (h messageHandler) Enabled() bool {
	return configs.AppConfig.Channels.SlackBotTestChannel.Messages.HiTestBotMessage.Enabled
}

func (h messageHandler) GetDescription() string {
	return configs.AppConfig.Channels.SlackBotTestChannel.Messages.HiTestBotMessage.Description
}

func (h messageHandler) GetText() string {
	return configs.AppConfig.Channels.SlackBotTestChannel.Messages.HiTestBotMessage.Text
}
