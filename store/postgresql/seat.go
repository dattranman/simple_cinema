package postgresql

import (
	"github.com/dattranman/simple_cinema/model/schema"
	"github.com/dattranman/simple_cinema/store"
)

type PostgresqlSeatStore struct {
	postgresql *PostgresStore
}

func NewSeatStore(postgresql *PostgresStore) store.Seat {
	return &PostgresqlSeatStore{postgresql: postgresql}
}

func (s *PostgresqlSeatStore) GetByRoomID(roomID int) ([]*schema.Seat, error) {
	var seats []*schema.Seat
	if err := s.postgresql.db.Where("room_id = ?", roomID).Find(&seats).Error; err != nil {
		return nil, err
	}
	return seats, nil
}

func (s *PostgresqlSeatStore) Create(seats []*schema.Seat) error {
	return s.postgresql.db.Create(seats).Error
}

func (s *PostgresqlSeatStore) Delete(roomID int, seats []*schema.Seat) error {
	tx := s.postgresql.db.Begin()
	for _, seat := range seats {
		err := tx.Where("room_id = ?", roomID).
			Where(`"row" = ?`, seat.Row).
			Where(`"column" = ?`, seat.Column).
			Delete(&schema.Seat{}).Error
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	return tx.Commit().Error
}
