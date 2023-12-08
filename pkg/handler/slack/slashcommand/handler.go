package slashcommand

import "github.com/slack-go/slack/socketmode"

type Handler interface {
	HandleCommand(evt *socketmode.Event, client *socketmode.Client) error
}
