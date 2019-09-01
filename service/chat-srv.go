package service

import (
	"chat/transport"
)

type ChatService struct {
	incoming []*transport.Event
}

func NewChatService() *ChatService {
	inc := make([]*transport.Event, 0)
	return &ChatService{incoming: inc}
}
