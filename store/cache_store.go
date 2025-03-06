package store

import "github.com/dattranman/simple_cinema/model/schema"

type CacheStore interface {
	Room() Room
}

type RoomCache interface {
	Get(id int) (*schema.Room, error)
	SetBookedSeat(id int, seats []*schema.Seat) error
	GetBookedSeats(roomDetail *schema.Room) ([]*schema.Seat, error)
	DeleteBookedSeat(id int, seats []*schema.Seat) error
}
