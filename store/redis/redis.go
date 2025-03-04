package redis

import (
	"context"

	"github.com/rs/zerolog/log"

	"github.com/dattranman/simple_cinema/model"
	"github.com/dattranman/simple_cinema/store"
	"github.com/redis/go-redis/v9"
)

type RedisCache struct {
	store.CacheStore
	client   *redis.Client
	settings model.CacheSetting

	room store.RoomCache
}

func NewRedis(settings model.CacheSetting) *RedisCache {
	p := &RedisCache{
		settings: settings,
	}

	p.initConnection()
	p.room = NewRoomCache(p)
	return p
}

func (p *RedisCache) initConnection() {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     p.settings.URI,
		Password: p.settings.Password,
		DB:       0,
	})
	_, err := redisClient.Ping(context.Background()).Result()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to open Redis connection")
	}
	p.client = redisClient
}

func (p *RedisCache) Room() store.RoomCache {
	return p.room
}
