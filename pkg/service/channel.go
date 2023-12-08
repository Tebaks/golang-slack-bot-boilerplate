package service

import "app/pkg/handler"

type Channel interface {
	GetID() string
	Enabled() bool
	GetMessages() []handler.MessageHandler
}
