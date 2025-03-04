package postgresql

import (
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/dattranman/simple_cinema/model/schema"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewMockGormDB(t *testing.T) (*sql.DB, *gorm.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create mock db: %v", err)
	}
	gorm, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to create mock gorm db: %v", err)
	}
	gorm.Debug()
	return db, gorm, mock
}
func TestPostgresqlSeatStore_Create(t *testing.T) {
	db, gormMock, mock := NewMockGormDB(t)
	defer db.Close()
	type fields struct {
		postgresql *PostgresStore
	}
	gormMock.Debug()
	type args struct {
		seats []*schema.Seat
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		expectSql string
		mock      sqlmock.Sqlmock
		wantErr   bool
	}{
		{
			name: "success case",
			fields: fields{
				postgresql: &PostgresStore{db: gormMock},
			},
			args: args{
				seats: []*schema.Seat{
					{RoomID: 1, Row: 1, Column: 1},
				},
			},
			expectSql: `^INSERT INTO "seats" (.+)$`,
			mock:      mock,
			wantErr:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &PostgresqlSeatStore{
				postgresql: tt.fields.postgresql,
			}
			// expectSql := regexp.QuoteMeta(tt.expectSql)
			tt.mock.ExpectBegin()
			tt.mock.ExpectExec(tt.expectSql).
				WillReturnResult(sqlmock.NewResult(1, 1))
			tt.mock.ExpectCommit()
			err := s.Create(tt.args.seats)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostgresqlSeatStore.Create() error = %v, wantErr %v", err, tt.wantErr)
			}

			// Ensure all expectations were met
			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("unfulfilled expectations: %s", err)
			}
		})
	}
}
