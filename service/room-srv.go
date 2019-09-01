package service

import (
	"chat/transport"
	"context"
)

type userRoomList struct {
	currentUser *transport.UserInfo
	roomList    []*transport.RoomInfo
}

func NewUserRoomList() *userRoomList {
	list := make([]*transport.RoomInfo, 1)
	return &userRoomList{roomList: list}
}

func (u userRoomList) GetRoomList(ctx context.Context, t *transport.Get) (*transport.RoomList, error) {
	return &transport.RoomList{
		RoomList: u.roomList,
	}, nil
}
