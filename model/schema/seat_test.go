package schema

import (
	"reflect"
	"testing"

	"github.com/dattranman/simple_cinema/model/response"
)

func TestSeat_ToResponse(t *testing.T) {
	type fields struct {
		RoomID int
		Row    int
		Column int
	}
	tests := []struct {
		name   string
		fields fields
		want   response.Seat
	}{
		{
			name: "test 1",
			fields: fields{
				RoomID: 1,
				Row:    1,
				Column: 1,
			},
			want: response.Seat{
				Row:    1,
				Column: 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Seat{
				RoomID: tt.fields.RoomID,
				Row:    tt.fields.Row,
				Column: tt.fields.Column,
			}
			if got := s.ToResponse(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Seat.ToResponse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSeatList_ToResponseList(t *testing.T) {
	tests := []struct {
		name string
		s    SeatList
		want []response.Seat
	}{
		{
			name: "test 1",
			s:    SeatList{Seat{RoomID: 1, Row: 1, Column: 1}},
			want: []response.Seat{response.Seat{Row: 1, Column: 1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.ToResponseList(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SeatList.ToResponseList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSeat_TableName(t *testing.T) {
	type fields struct {
		RoomID int
		Row    int
		Column int
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "test 1",
			fields: fields{
				RoomID: 1,
				Row:    1,
				Column: 1,
			},
			want: "seats",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Seat{
				RoomID: tt.fields.RoomID,
				Row:    tt.fields.Row,
				Column: tt.fields.Column,
			}
			if got := s.TableName(); got != tt.want {
				t.Errorf("Seat.TableName() = %v, want %v", got, tt.want)
			}
		})
	}
}
