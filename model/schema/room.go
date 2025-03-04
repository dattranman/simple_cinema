package schema

import (
	"github.com/dattranman/simple_cinema/model/response"
)

const (
	RoomTable = "rooms"
)

type Room struct {
	ID          int `json:"id"`
	Row         int `json:"row"`
	Column      int `json:"column"`
	MinDistance int `json:"min_distance"`
}

func (r *Room) TableName() string {
	return RoomTable
}

func (r *Room) ToResponse() *response.Room {
	return &response.Room{
		ID:          r.ID,
		Row:         r.Row,
		Column:      r.Column,
		MinDistance: r.MinDistance,
	}
}
