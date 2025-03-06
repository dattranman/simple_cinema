package schema

import (
	"reflect"
	"testing"

	"github.com/dattranman/simple_cinema/model/response"
)

func TestRoom_ToResponse(t *testing.T) {
	type fields struct {
		ID          int
		Row         int
		Column      int
		MinDistance int
	}
	tests := []struct {
		name   string
		fields fields
		want   *response.Room
	}{
		{
			name: "test 1",
			fields: fields{
				ID:          1,
				Row:         10,
				Column:      10,
				MinDistance: 1,
			},
			want: &response.Room{
				ID:          1,
				Row:         10,
				Column:      10,
				MinDistance: 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Room{
				ID:          tt.fields.ID,
				Row:         tt.fields.Row,
				Column:      tt.fields.Column,
				MinDistance: tt.fields.MinDistance,
			}
			if got := r.ToResponse(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Room.ToResponse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRoom_TableName(t *testing.T) {
	type fields struct {
		ID          int
		Row         int
		Column      int
		MinDistance int
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "test 1",
			fields: fields{
				ID:          1,
				Row:         10,
				Column:      10,
				MinDistance: 1,
			},
			want: "rooms",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Room{
				ID:          tt.fields.ID,
				Row:         tt.fields.Row,
				Column:      tt.fields.Column,
				MinDistance: tt.fields.MinDistance,
			}
			if got := r.TableName(); got != tt.want {
				t.Errorf("Room.TableName() = %v, want %v", got, tt.want)
			}
		})
	}
}
