package eventsapievent

import (
	"app/pkg/service"
	"fmt"
	"github.com/slack-go/slack/slackevents"
	"github.com/slack-go/slack/socketmode"
	"strings"
)

type Handler interface {
	Handle(evt *socketmode.Event, client *socketmode.Client, channels []service.Channel) error
}

type handler struct {
}

func NewHandler() Handler {
	return &handler{}
}

func (h *handler) Handle(evt *socketmode.Event, client *socketmode.Client, channels []service.Channel) error {
	eventsAPIEvent, ok := evt.Data.(slackevents.EventsAPIEvent)
	if !ok {
		return fmt.Errorf("unexpected event type: %s", evt.Type)
	}

	if eventsAPIEvent.Type == slackevents.CallbackEvent {
		innerEvent := eventsAPIEvent.InnerEvent

		switch ev := innerEvent.Data.(type) {
		case *slackevents.MessageEvent:
			for _, channel := range channels {
				if !channel.Enabled() {
					continue
				}
				if ev.Channel != channel.GetID() {
					continue
				}
				for _, messageHandler := range channel.GetMessages() {
					if !messageHandler.Enabled() {
						continue
					}
					if strings.Contains(ev.Text, messageHandler.GetText()) {
						err := messageHandler.HandleMessage(ev, client)
						if err != nil {
							fmt.Printf("Error while handling message: %+v\n", err)
							continue
						}
					}
				}
			}
		}
	}
	client.Ack(*evt.Request)
	return nil
}
