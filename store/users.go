package store

import (
	"chat/transport"
	"log"
)

type UserStore struct {
	users []transport.RegisterInfo
}

func NewUserStore() *UserStore {
	users := make([]transport.RegisterInfo, 0)
	return &UserStore{users: users}
}

func (s UserStore) CheckUserExist(t *transport.UserInfo) bool {
	for _, v := range s.users {
		if v.UserInfo.Name == t.Name {
			return true
		}
	}
	return false
}

func (s *UserStore) RegisterUser(t *transport.UserInfo) (ok bool, err error) {
	NewId, err := CalcHash(t.XXX_unrecognized)
	if err != nil {
		log.Println(err)
		return false, err
	}
	ui := transport.UserInfo{
		Id:   NewId,
		Name: t.Name,
	}
	rl := transport.RoomList{
		Id:       NewId,
		RoomList: make([]*transport.RoomInfo, 0),
	}
	regi := transport.RegisterInfo{
		RoomList: &rl,
		UserInfo: &ui,
	}
	s.users = append(s.users, regi)
	return true, err
}
