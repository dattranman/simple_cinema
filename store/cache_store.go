package store

import "github.com/dattranman/simple_cinema/model/schema"

type CacheStore interface {
	Room() Room
}

type RoomCache interface {
	Get(id string) (*schema.Room, error)
	SetBookedSeat(id string, seats []schema.Seat) error
	GetBookedSeat(id string) ([]schema.Seat, error)
	DeleteBookedSeat(id string, seats []schema.Seat) error
}
