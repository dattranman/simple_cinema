package postgresql

import (
	"database/sql"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/dattranman/simple_cinema/model/schema"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewMockGormDB(t *testing.T) (*sql.DB, *gorm.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create mock db: %v", err)
	}
	gorm, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{
		QueryFields: true,
		PrepareStmt: true,
	})
	if err != nil {
		t.Fatalf("failed to create mock gorm db: %v", err)
	}
	gorm.Debug()
	return db, gorm, mock
}
func TestPostgresqlSeatStore_Create(t *testing.T) {
	type args struct {
		seats []*schema.Seat
	}
	tests := []struct {
		name      string
		args      args
		expectSql string
		wantErr   bool
	}{
		{
			name: "success case - single seat",
			args: args{
				seats: []*schema.Seat{
					{RoomID: 1, Row: 1, Column: 1},
				},
			},
			expectSql: `INSERT INTO "seats" ("room_id","row","column") VALUES ($1,$2,$3)`,
			wantErr:   false,
		}, {
			name: "success case - multiple seats",
			args: args{
				seats: []*schema.Seat{
					{RoomID: 1, Row: 1, Column: 1},
					{RoomID: 1, Row: 1, Column: 2},
				},
			},
			expectSql: `INSERT INTO "seats" ("room_id","row","column") VALUES ($1,$2,$3),($4,$5,$6)`,
			wantErr:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, gormMock, mock := NewMockGormDB(t)
			defer db.Close()
			gormMock.Debug()
			s := &PostgresqlSeatStore{
				postgresql: &PostgresStore{db: gormMock},
			}

			if len(tt.args.seats) > 0 {
				expectSql := regexp.QuoteMeta(tt.expectSql)
				mock.ExpectBegin()
				mock.ExpectPrepare(expectSql)
				mock.ExpectPrepare(expectSql)
				mock.ExpectExec(expectSql).
					WillReturnResult(sqlmock.NewResult(1, int64(len(tt.args.seats))))
				mock.ExpectCommit()
			}

			err := s.Create(tt.args.seats)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			// Ensure all expectations were met
			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("unfulfilled expectations: %s", err)
			}
		})
	}
}

func TestPostgresqlSeatStore_Delete(t *testing.T) {
	type args struct {
		roomID int
		seats  []*schema.Seat
	}
	tests := []struct {
		name      string
		args      args
		expectSql string
		mock      sqlmock.Sqlmock
		wantErr   bool
	}{
		{
			name: "success case",
			args: args{
				roomID: 1,
				seats: []*schema.Seat{
					{RoomID: 1, Row: 1, Column: 1},
				},
			},
			expectSql: `DELETE FROM "seats" WHERE room_id = $1 AND "row" = $2 AND "column" = $3`,
			wantErr:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, gormMock, mock := NewMockGormDB(t)
			defer db.Close()
			s := &PostgresqlSeatStore{
				postgresql: &PostgresStore{db: gormMock},
			}
			expectSql := regexp.QuoteMeta(tt.expectSql)
			mock.ExpectBegin()
			mock.ExpectPrepare(expectSql)
			mock.ExpectPrepare(expectSql)
			mock.ExpectExec(expectSql).
				WillReturnResult(sqlmock.NewResult(1, 1))
			mock.ExpectCommit()
			err := s.Delete(tt.args.roomID, tt.args.seats)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("unfulfilled expectations: %s", err)
			}
		})
	}
}

func TestPostgresqlSeatStore_GetByRoomID(t *testing.T) {
	type args struct {
		roomID int
	}
	tests := []struct {
		name      string
		args      args
		expectSql string
		mock      sqlmock.Sqlmock
		rows      *sqlmock.Rows
		want      []*schema.Seat
		wantErr   bool
	}{
		{
			name: "success case more than one row",
			args: args{
				roomID: 1,
			},
			expectSql: `SELECT "seats"."room_id","seats"."row","seats"."column" FROM "seats" WHERE room_id = $1`,
			rows:      sqlmock.NewRows([]string{"room_id", "row", "column"}).AddRow(1, 1, 1).AddRow(1, 2, 2),
			want:      []*schema.Seat{{RoomID: 1, Row: 1, Column: 1}, {RoomID: 1, Row: 2, Column: 2}},
			wantErr:   false,
		}, {
			name: "not found",
			args: args{
				roomID: 1,
			},
			expectSql: `SELECT "seats"."room_id","seats"."row","seats"."column" FROM "seats" WHERE room_id = $1`,
			rows:      sqlmock.NewRows([]string{"room_id", "row", "column"}),
			want:      []*schema.Seat{},
			wantErr:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, gormMock, mock := NewMockGormDB(t)
			defer db.Close()
			s := &PostgresqlSeatStore{
				postgresql: &PostgresStore{db: gormMock},
			}
			expectSql := regexp.QuoteMeta(tt.expectSql)
			mock.ExpectPrepare(expectSql)
			if tt.rows != nil {
				mock.ExpectQuery(expectSql).
					WillReturnRows(tt.rows)
			}
			if tt.rows == nil {
				mock.ExpectQuery(tt.expectSql).
					WillReturnError(sql.ErrNoRows)
			}
			got, err := s.GetByRoomID(tt.args.roomID)
			if tt.wantErr {
				assert.Error(t, err)
			}
			if !tt.wantErr {
				assert.NoError(t, err)
			}
			assert.Equal(t, tt.want, got)
		})
	}
}
