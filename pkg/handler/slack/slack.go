package slack

import (
	"app/pkg/handler/slack/eventsapievent"
	"app/pkg/handler/slack/slashcommandevent"
	"app/pkg/service"
	"github.com/slack-go/slack/socketmode"
)

type Handler interface {
	Handle(client *socketmode.Client, channels []service.Channel)
}

type handler struct {
	slashCommandEventHandler slashcommandevent.Handler
	eventAPIEventHandler     eventsapievent.Handler
}

func NewHandler(slashCommandHandler slashcommandevent.Handler,
	eventsAPIHandler eventsapievent.Handler) Handler {
	return handler{
		slashCommandEventHandler: slashCommandHandler,
		eventAPIEventHandler:     eventsAPIHandler,
	}
}

func (h handler) Handle(client *socketmode.Client, channels []service.Channel) {
	go func() {
		for evt := range client.Events {
			switch evt.Type {
			case socketmode.EventTypeSlashCommand:
				err := h.slashCommandEventHandler.Handle(&evt, client)
				if err != nil {
					return
				}
			case socketmode.EventTypeEventsAPI:
				err := h.eventAPIEventHandler.Handle(&evt, client, channels)
				if err != nil {
					return
				}
			}
		}
	}()
}
