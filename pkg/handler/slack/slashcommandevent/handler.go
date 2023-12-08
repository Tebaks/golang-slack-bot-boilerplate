package slashcommandevent

import (
	"app/pkg/handler/slack/slashcommand"
	"github.com/slack-go/slack"
	"github.com/slack-go/slack/socketmode"
)

type Handler interface {
	Handle(evt *socketmode.Event, client *socketmode.Client) error
}

type handler struct {
	sayHiHandler slashcommand.Handler
}

func NewHandler(sayHiCommandHandler slashcommand.Handler) Handler {
	return &handler{
		sayHiHandler: sayHiCommandHandler,
	}
}

func (h *handler) Handle(evt *socketmode.Event, client *socketmode.Client) error {
	data := evt.Data.(slack.SlashCommand)
	switch data.Command {
	case "/hi":
		err := h.sayHiHandler.HandleCommand(evt, client)
		if err != nil {
			return err
		}
	}

	client.Ack(*evt.Request)
	return nil
}
