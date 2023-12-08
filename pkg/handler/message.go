package handler

import (
	"github.com/slack-go/slack/slackevents"
	"github.com/slack-go/slack/socketmode"
)

type MessageHandler interface {
	HandleMessage(messageEvent *slackevents.MessageEvent, client *socketmode.Client) error
	Enabled() bool
	GetDescription() string
	GetText() string
}
