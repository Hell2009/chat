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

func (s *server) CreateUser(ctx context.Context, t *transport.UserInfo) (*transport.Status, error) {
	userExist := s.userList.CheckUserExist(t)
	if !userExist {
		_, err := s.userList.RegisterUser(t)
		if err != nil {
			return &transport.Status{
				Ok: false,
			}, err
		}
	}
	return &transport.Status{
		Ok: true,
	}, nil
}

func (s server) LogInUser(ctx context.Context, t *transport.UserInfo) (*transport.Status, error) {
	for ok := s.userList.CheckUserExist(t); ok; {
		return &transport.Status{
			Ok: true,
		}, nil
	}
	return &transport.Status{
		Ok: false,
	}, errors.New("user not exist")
}
