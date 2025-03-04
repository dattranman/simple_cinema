package redis

import (
	"context"

	"github.com/dattranman/simple_cinema/model/schema"
	"github.com/dattranman/simple_cinema/store"
)

type RedisRoomCache struct {
	redis *RedisCache
}

func NewRoomCache(redis *RedisCache) store.RoomCache {
	return &RedisRoomCache{
		redis: redis,
	}
}

func (r *RedisRoomCache) Get(id string) (*schema.Room, error) {
	return nil, nil
}

func (r *RedisRoomCache) SetBookedSeat(id string, seats []schema.Seat) error {
	for _, seat := range seats {
		r.redis.client.SetBit(context.Background(), id, int64(seat.Row), 1)
	}
	return nil
}

func (r *RedisRoomCache) GetBookedSeat(id string) ([]schema.Seat, error) {
	seats := []schema.Seat{}
	for i := 0; i < 10; i++ {
		if r.redis.client.GetBit(context.Background(), id, int64(i)).Val() == 1 {
			seats = append(seats, schema.Seat{Row: i})
		}
	}
	return seats, nil
}

func (r *RedisRoomCache) DeleteBookedSeat(id string, seats []schema.Seat) error {
	for _, seat := range seats {
		r.redis.client.SetBit(context.Background(), id, int64(seat.Row), 0)
	}
	return nil
}
