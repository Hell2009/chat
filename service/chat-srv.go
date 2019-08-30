package service

import (
	"context"
	"exp-chat/transport"
)

type ChatService struct {
	incoming []*transport.Event
}

func NewChatService() *ChatService {
	inc := make([]*transport.Event, 0)
	return &ChatService{incoming: inc}
}

func (c *ChatService) Push(ctx context.Context, t *transport.Event) (*transport.Status, error) {
	panic("implement me")
}

func (c ChatService) SendMessage(ctx context.Context, t *transport.Msg) (*transport.Status, error) {
	panic("implement me")
}

func (c ChatService) Pull(t *transport.Get, stream transport.Chat_PullServer) error {
	panic("implement me")
}
