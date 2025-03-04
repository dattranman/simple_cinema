package schema

import "github.com/dattranman/simple_cinema/model/response"

const (
	seatTableName = "seats"
)

type Seat struct {
	RoomID int `json:"room_id"`
	Row    int `json:"row"`
	Column int `json:"column"`
}

func (s *Seat) TableName() string {
	return seatTableName
}

func (s *Seat) ToResponse() response.Seat {
	return response.Seat{
		Row:    uint(s.Row),
		Column: uint(s.Column),
	}
}
