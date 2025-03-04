package postgresql

import (
	"github.com/dattranman/simple_cinema/model/schema"
	"github.com/dattranman/simple_cinema/store"
)

type PostgresRoomStore struct {
	postgres *PostgresStore
}

func NewRoomStore(pg *PostgresStore) store.Room {
	return &PostgresRoomStore{
		postgres: pg,
	}
}

func (r *PostgresRoomStore) GetList() (rooms []*schema.Room, total int64, err error) {
	err = r.postgres.db.Find(&rooms).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}
	return rooms, total, nil
}

func (r *PostgresRoomStore) Create(room *schema.Room) error {
	return r.postgres.db.Create(room).Error
}

func (r *PostgresRoomStore) Update(room *schema.Room) error {
	return r.postgres.db.Save(room).Error
}

func (r *PostgresRoomStore) Delete(id int) error {
	return r.postgres.db.Delete(&schema.Room{}, "id = ?", id).Error
}

func (r *PostgresRoomStore) GetByID(id int) (*schema.Room, error) {
	var room *schema.Room
	err := r.postgres.db.Where("id = ?", id).First(&room).Error
	if err != nil {
		return nil, err
	}
	return room, nil
}
