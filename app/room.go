package app

import (
	"net/http"

	"github.com/dattranman/simple_cinema/model/request"
	"github.com/dattranman/simple_cinema/model/response"
	"github.com/dattranman/simple_cinema/model/schema"
)

func (a *App) GetRoomList() (resp *response.GetRoomList, err error) {
	rooms, _, err := a.Store.Room().GetList()
	if err != nil {
		return nil, err
	}
	roomsResponse := make([]*response.Room, len(rooms))
	for i, room := range rooms {
		roomsResponse[i] = room.ToResponse()
	}
	resp = &response.GetRoomList{
		Base: response.Base{
			Code:    http.StatusOK,
			Message: "Success",
		},
		Data:  roomsResponse,
		Total: len(rooms),
	}
	return resp, nil
}

func (a *App) GetRoomDetail(id int) (resp *response.GetRoomDetail, err error) {
	room, err := a.Store.Room().GetByID(id)
	if err != nil {
		return nil, err
	}
	resp = &response.GetRoomDetail{
		Base: response.Base{
			Code:    http.StatusOK,
			Message: "Success",
		},
		Data: room.ToResponse(),
	}
	return resp, nil
}

func (a *App) CreateRoom(req request.CreateRoom) (resp *response.CreateRoom, err error) {
	room := &schema.Room{
		Row:         int(req.Row),
		Column:      int(req.Column),
		MinDistance: int(req.MinDistance),
	}
	err = a.Store.Room().Create(room)
	if err != nil {
		return nil, err
	}
	resp = &response.CreateRoom{
		Base: response.Base{
			Code:    http.StatusOK,
			Message: "Success",
		},
		Data: room.ToResponse(),
	}
	return resp, nil
}

func (a *App) DeleteRoom(id int) (err error) {
	err = a.Store.Room().Delete(id)
	if err != nil {
		return err
	}
	return nil
}
