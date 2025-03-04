package app

import (
	"github.com/dattranman/simple_cinema/model"
	"github.com/dattranman/simple_cinema/store/postgresql"
	"github.com/dattranman/simple_cinema/store/redis"
	"github.com/rs/zerolog/log"
)

type App struct {
	Config *model.Configuration

	Store *postgresql.PostgresStore
	Cache *redis.RedisCache
}

func New(cfg *model.Configuration) (*App, error) {
	app := &App{
		Config: cfg,
	}

	log.Info().Msg("Server is initializing...")
	app.Store = postgresql.NewPostgres(app.Config.SQLSettings)
	app.Cache = redis.NewRedis(app.Config.CacheSettings)

	return app, nil
}
