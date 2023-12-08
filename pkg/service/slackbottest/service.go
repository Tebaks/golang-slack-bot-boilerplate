package slackbottest

import (
	"app/pkg/configs"
	"app/pkg/handler"
)

type Service struct {
	hiTestBotMessageHandler handler.MessageHandler
}

func NewService(hiTestBotMessageHandler handler.MessageHandler) *Service {
	return &Service{
		hiTestBotMessageHandler: hiTestBotMessageHandler,
	}
}

func (s *Service) GetID() string {
	return configs.AppConfig.Channels.SlackBotTestChannel.ID
}

func (s *Service) Enabled() bool {
	return configs.AppConfig.Channels.SlackBotTestChannel.Enabled
}

func (s *Service) GetMessages() []handler.MessageHandler {
	return []handler.MessageHandler{s.hiTestBotMessageHandler}
}
