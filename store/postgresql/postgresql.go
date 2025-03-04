package postgresql

import (
	"github.com/rs/zerolog/log"

	"github.com/dattranman/simple_cinema/model"
	"github.com/dattranman/simple_cinema/store"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type PostgresStore struct {
	store.SqlStore
	settings model.SQLSettings
	db       *gorm.DB

	room store.Room
	seat store.Seat
}

func NewPostgres(settings model.SQLSettings) *PostgresStore {
	p := &PostgresStore{
		settings: settings,
	}

	p.initConnection()
	p.room = NewRoomStore(p)
	p.seat = NewSeatStore(p)
	return p
}

func (p *PostgresStore) initConnection() {
	config := gorm.Config{
		QueryFields: true,
		PrepareStmt: true,
	}
	if p.settings.Debug {
		config.Logger = logger.Default.LogMode(logger.Info)
	}
	pg := postgres.New(postgres.Config{
		DSN:                  p.settings.URI,
		PreferSimpleProtocol: true,
	})
	db, err := gorm.Open(pg, &config)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to open SQL connection")
	}
	p.db = db
}

func (p *PostgresStore) Room() store.Room {
	return p.room
}

func (p *PostgresStore) Seat() store.Seat {
	return p.seat
}
