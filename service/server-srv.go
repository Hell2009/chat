package service

import (
	"chat/store"
	"chat/transport"
	"context"
	"errors"
)

type server struct {
	userList store.UserStore
}

func NewServer(userList store.UserStore) *server {
	return &server{userList: userList}
}

func (s *server) CreateUser(ctx context.Context, t *transport.UserInfo) (*transport.RegisterInfo, error) {
	userExist := s.userList.CheckUserExist(t)
	if !userExist {
		_, err := s.userList.RegisterUser(t)
		if err != nil {
			return &transport.RegisterInfo{}, err
		}
	}
	return &transport.RegisterInfo{}, nil
}

func (s server) LogInUser(ctx context.Context, t *transport.UserInfo) (*transport.RegisterInfo, error) {
	for ok := s.userList.CheckUserExist(t); ok; {
		return &transport.RegisterInfo{}, nil
	}
	return &transport.RegisterInfo{}, errors.New("user not exist")
}

func (s server) PullEvent(*transport.UserInfo, transport.Server_PullEventServer) error {
	panic("implement me")
}
