package store

import (
	"chat/transport"
	"errors"
	"log"
)

type RoomStore struct {
	rooms []*transport.RoomInfo
}

func NewRoomStore() *RoomStore {
	rooms := make([]*transport.RoomInfo, 0)
	return &RoomStore{rooms: rooms}
}

func (s RoomStore) CheckRoomExist(t *transport.RoomInfo) (bool, error) {
	for _, v := range s.rooms {
		if v.Id == t.Id {
			return true, nil
		}
	}
	return false, errors.New("room not exist")
}

func (s *RoomStore) RegisterRoom(t *transport.RoomInfo) (ok bool, err error) {
	NewId, err := CalcHash(t.XXX_unrecognized)
	if err != nil {
		log.Println(err)
		return false, err
	}
	ri := transport.RoomInfo{
		Id:       NewId,
		Name:     t.Name,
		Owner:    t.Owner,
		ListUser: t.ListUser,
	}
	s.rooms = append(s.rooms, &ri)
	return true, err
}

func CheckUserInRoom(u *transport.UserInfo, lu []*transport.UserInfo) (int, error) {
	for i, v := range lu {
		if v == u {
			return i, nil
		}
	}
	return 0, errors.New("user not found")
}

func (s *RoomStore) KickUser(t *transport.RoomUser) (bool, error) {
	_, err := s.CheckRoomExist(t.Room)
	if err != nil {
		return false, err
	}
	userIndex, err := CheckUserInRoom(t.User, t.Room.ListUser)
	if err != nil {
		log.Println(err)
		return false, err
	}
	t.Room.ListUser = append(t.Room.ListUser[:userIndex], t.Room.ListUser[userIndex+1:]...)
	return true, nil
}
