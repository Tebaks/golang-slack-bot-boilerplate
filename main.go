package main

import (
	"app/pkg/configs"
	"app/pkg/handler/hitestbot"
	slackhandler "app/pkg/handler/slack"
	"app/pkg/handler/slack/eventsapievent"
	"app/pkg/handler/slack/slashcommand/sayhi"
	"app/pkg/handler/slack/slashcommandevent"
	"app/pkg/service"
	"app/pkg/service/slackbottest"
	"fmt"
	"github.com/slack-go/slack"
	"github.com/slack-go/slack/socketmode"
	"log"
	"os"
)

func main() {
	configs.Init()

	token := configs.AppCredentials.SlackCredentials.SlackAuthToken
	appToken := configs.AppCredentials.SlackCredentials.SlackAppToken

	client := slack.New(token, slack.OptionDebug(true), slack.OptionAppLevelToken(appToken))

	app := socketmode.New(
		client,
		socketmode.OptionDebug(true),
		socketmode.OptionLog(log.New(os.Stdout, "socketmode: ", log.Lshortfile|log.LstdFlags)),
	)

	sayHiCommandHandler := sayhi.NewHandler()
	slashCommandEventHandler := slashcommandevent.NewHandler(sayHiCommandHandler)
	eventsAPIEventHandler := eventsapievent.NewHandler()
	hiTestBotMessageHandler := hitestbot.NewMessageHandler()

	slackBotTestService := slackbottest.NewService(hiTestBotMessageHandler)
	channels := []service.Channel{slackBotTestService}

	slackHandler := slackhandler.NewHandler(slashCommandEventHandler, eventsAPIEventHandler)
	slackHandler.Handle(app, channels)

	err := app.Run()
	if err != nil {
		fmt.Println(err)
		return
	}
}
