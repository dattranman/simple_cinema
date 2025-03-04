package store

import (
	"github.com/dattranman/simple_cinema/model/schema"
)

type SqlStore interface {
	Room() Room
}

type Room interface {
	Create(room *schema.Room) error
	GetByID(id int) (*schema.Room, error)
	GetList() ([]*schema.Room, int64, error)
	Update(room *schema.Room) error
	Delete(id int) error
}

type Seat interface {
	GetByRoomID(roomID int) ([]*schema.Seat, error)
	Create(seats []*schema.Seat) error
	Delete(roomID int, seats []*schema.Seat) error
}
