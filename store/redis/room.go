package redis

import (
	"context"
	"fmt"
	"hash/crc32"

	"github.com/dattranman/simple_cinema/model/schema"
	"github.com/dattranman/simple_cinema/store"
)

const (
	BookedSeatKey = "room:%d:booked_seat"
	seatFormat    = "%d-%d"
)

type RedisRoomCache struct {
	redis *RedisCache
}

func NewRoomCache(redis *RedisCache) store.RoomCache {
	return &RedisRoomCache{
		redis: redis,
	}
}

func (r *RedisRoomCache) Get(id int) (*schema.Room, error) {
	return nil, nil
}

func hashKeyByCrc32(key string) int {
	hash := crc32.NewIEEE()
	hash.Write([]byte(key))
	return int(hash.Sum32())
}

func (r *RedisRoomCache) SetBookedSeat(id int, seats []*schema.Seat) error {
	key := fmt.Sprintf(BookedSeatKey, id)
	for _, seat := range seats {
		seatKey := fmt.Sprintf(seatFormat, seat.Row, seat.Column)
		seatKeyHash := hashKeyByCrc32(seatKey)
		r.redis.client.SetBit(context.Background(), key, int64(seatKeyHash), 1)
	}
	return nil
}

func (r *RedisRoomCache) GetBookedSeats(roomDetail *schema.Room) ([]*schema.Seat, error) {
	key := fmt.Sprintf(BookedSeatKey, roomDetail.ID)
	seats := []*schema.Seat{}
	for i := 0; i < roomDetail.Row; i++ {
		for j := 0; j < roomDetail.Column; j++ {
			seatKey := fmt.Sprintf(seatFormat, i, j)
			seatKeyHash := hashKeyByCrc32(seatKey)
			bit := r.redis.client.GetBit(context.Background(), key, int64(seatKeyHash))
			if bit.Val() == 1 {
				seats = append(seats, &schema.Seat{Row: i, Column: j})
			}
		}
	}
	return seats, nil
}

func (r *RedisRoomCache) DeleteBookedSeat(id int, seats []*schema.Seat) error {
	key := fmt.Sprintf(BookedSeatKey, id)
	for _, seat := range seats {
		seatKey := fmt.Sprintf(seatFormat, seat.Row, seat.Column)
		seatKeyHash := hashKeyByCrc32(seatKey)
		r.redis.client.SetBit(context.Background(), key, int64(seatKeyHash), 0)
	}
	return nil
}
