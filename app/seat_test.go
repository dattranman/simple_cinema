package app

import (
	"reflect"
	"testing"

	"github.com/dattranman/simple_cinema/model/schema"
)

func TestCalculateMinDistanceFromTwoGroupSeat(t *testing.T) {
	type args struct {
		group1 []*schema.Seat
		group2 []*schema.Seat
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test 1",
			args: args{
				group1: []*schema.Seat{{Row: 1, Column: 1}, {Row: 1, Column: 2}},
				group2: []*schema.Seat{{Row: 2, Column: 1}, {Row: 2, Column: 2}},
			},
			want: 1,
		},
		{
			name: "test 2",
			args: args{
				group1: []*schema.Seat{{Row: 1, Column: 1}, {Row: 1, Column: 2}, {Row: 15, Column: 14}},
				group2: []*schema.Seat{{Row: 8, Column: 2}, {Row: 8, Column: 3}},
			},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalculateMinDistanceFromTwoGroupSeat(tt.args.group1, tt.args.group2); got != tt.want {
				t.Errorf("CalculateMinDistanceFromTwoGroupSeat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetAvailableSeat(t *testing.T) {
	type args struct {
		bookedSeat  []*schema.Seat
		Room        *schema.Room
		minDistance int
	}
	tests := []struct {
		name string
		args args
		want []*schema.Seat
	}{
		{
			name: "test 1",
			args: args{
				bookedSeat: []*schema.Seat{{Row: 1, Column: 1}, {Row: 1, Column: 2}},
				Room:       &schema.Room{Row: 10, Column: 10, MinDistance: 1},
			},
			want: []*schema.Seat{
				{Row: 0, Column: 0}, {Row: 0, Column: 1}, {Row: 0, Column: 2}, {Row: 0, Column: 3}, {Row: 0, Column: 4},
				{Row: 0, Column: 5}, {Row: 0, Column: 6}, {Row: 0, Column: 7}, {Row: 0, Column: 8}, {Row: 0, Column: 9},
				{Row: 1, Column: 0}, {Row: 1, Column: 3}, {Row: 1, Column: 4}, {Row: 1, Column: 5}, {Row: 1, Column: 6},
				{Row: 1, Column: 7}, {Row: 1, Column: 8}, {Row: 1, Column: 9}, {Row: 2, Column: 0}, {Row: 2, Column: 1},
				{Row: 2, Column: 2}, {Row: 2, Column: 3}, {Row: 2, Column: 4}, {Row: 2, Column: 5}, {Row: 2, Column: 6},
				{Row: 2, Column: 7}, {Row: 2, Column: 8}, {Row: 2, Column: 9}, {Row: 3, Column: 0}, {Row: 3, Column: 1},
				{Row: 3, Column: 2}, {Row: 3, Column: 3}, {Row: 3, Column: 4}, {Row: 3, Column: 5}, {Row: 3, Column: 6},
				{Row: 3, Column: 7}, {Row: 3, Column: 8}, {Row: 3, Column: 9}, {Row: 4, Column: 0}, {Row: 4, Column: 1},
				{Row: 4, Column: 2}, {Row: 4, Column: 3}, {Row: 4, Column: 4}, {Row: 4, Column: 5}, {Row: 4, Column: 6},
				{Row: 4, Column: 7}, {Row: 4, Column: 8}, {Row: 4, Column: 9}, {Row: 5, Column: 0}, {Row: 5, Column: 1},
				{Row: 5, Column: 2}, {Row: 5, Column: 3}, {Row: 5, Column: 4}, {Row: 5, Column: 5}, {Row: 5, Column: 6},
				{Row: 5, Column: 7}, {Row: 5, Column: 8}, {Row: 5, Column: 9}, {Row: 6, Column: 0}, {Row: 6, Column: 1},
				{Row: 6, Column: 2}, {Row: 6, Column: 3}, {Row: 6, Column: 4}, {Row: 6, Column: 5}, {Row: 6, Column: 6},
				{Row: 6, Column: 7}, {Row: 6, Column: 8}, {Row: 6, Column: 9}, {Row: 7, Column: 0}, {Row: 7, Column: 1},
				{Row: 7, Column: 2}, {Row: 7, Column: 3}, {Row: 7, Column: 4}, {Row: 7, Column: 5}, {Row: 7, Column: 6},
				{Row: 7, Column: 7}, {Row: 7, Column: 8}, {Row: 7, Column: 9}, {Row: 8, Column: 0}, {Row: 8, Column: 1},
				{Row: 8, Column: 2}, {Row: 8, Column: 3}, {Row: 8, Column: 4}, {Row: 8, Column: 5}, {Row: 8, Column: 6},
				{Row: 8, Column: 7}, {Row: 8, Column: 8}, {Row: 8, Column: 9}, {Row: 9, Column: 0}, {Row: 9, Column: 1},
				{Row: 9, Column: 2}, {Row: 9, Column: 3}, {Row: 9, Column: 4}, {Row: 9, Column: 5}, {Row: 9, Column: 6},
				{Row: 9, Column: 7}, {Row: 9, Column: 8}, {Row: 9, Column: 9},
			},
		}, {
			name: "test 2",
			args: args{
				bookedSeat: []*schema.Seat{{Row: 1, Column: 1}, {Row: 1, Column: 2}},
				Room:       &schema.Room{Row: 10, Column: 10, MinDistance: 3},
			},
			want: []*schema.Seat{
				{Row: 0, Column: 4}, {Row: 0, Column: 5}, {Row: 0, Column: 6}, {Row: 0, Column: 7}, {Row: 0, Column: 8}, {Row: 0, Column: 9},
				{Row: 1, Column: 5}, {Row: 1, Column: 6}, {Row: 1, Column: 7}, {Row: 1, Column: 8}, {Row: 1, Column: 9},
				{Row: 2, Column: 4}, {Row: 2, Column: 5}, {Row: 2, Column: 6}, {Row: 2, Column: 7}, {Row: 2, Column: 8}, {Row: 2, Column: 9},
				{Row: 3, Column: 0}, {Row: 3, Column: 3}, {Row: 3, Column: 4}, {Row: 3, Column: 5}, {Row: 3, Column: 6}, {Row: 3, Column: 7}, {Row: 3, Column: 8}, {Row: 3, Column: 9},
				{Row: 4, Column: 0}, {Row: 4, Column: 1}, {Row: 4, Column: 2}, {Row: 4, Column: 3}, {Row: 4, Column: 4}, {Row: 4, Column: 5}, {Row: 4, Column: 6}, {Row: 4, Column: 7}, {Row: 4, Column: 8}, {Row: 4, Column: 9},
				{Row: 5, Column: 0}, {Row: 5, Column: 1}, {Row: 5, Column: 2}, {Row: 5, Column: 3}, {Row: 5, Column: 4}, {Row: 5, Column: 5}, {Row: 5, Column: 6}, {Row: 5, Column: 7}, {Row: 5, Column: 8}, {Row: 5, Column: 9},
				{Row: 6, Column: 0}, {Row: 6, Column: 1}, {Row: 6, Column: 2}, {Row: 6, Column: 3}, {Row: 6, Column: 4}, {Row: 6, Column: 5}, {Row: 6, Column: 6}, {Row: 6, Column: 7}, {Row: 6, Column: 8}, {Row: 6, Column: 9},
				{Row: 7, Column: 0}, {Row: 7, Column: 1}, {Row: 7, Column: 2}, {Row: 7, Column: 3}, {Row: 7, Column: 4}, {Row: 7, Column: 5}, {Row: 7, Column: 6}, {Row: 7, Column: 7}, {Row: 7, Column: 8}, {Row: 7, Column: 9},
				{Row: 8, Column: 0}, {Row: 8, Column: 1}, {Row: 8, Column: 2}, {Row: 8, Column: 3}, {Row: 8, Column: 4}, {Row: 8, Column: 5}, {Row: 8, Column: 6}, {Row: 8, Column: 7}, {Row: 8, Column: 8}, {Row: 8, Column: 9},
				{Row: 9, Column: 0}, {Row: 9, Column: 1}, {Row: 9, Column: 2}, {Row: 9, Column: 3}, {Row: 9, Column: 4}, {Row: 9, Column: 5}, {Row: 9, Column: 6}, {Row: 9, Column: 7}, {Row: 9, Column: 8}, {Row: 9, Column: 9},
			},
		}, {
			name: "test 3",
			args: args{
				bookedSeat: []*schema.Seat{},
				Room:       &schema.Room{Row: 4, Column: 5, MinDistance: 6},
			},
			want: []*schema.Seat{
				{Row: 0, Column: 0}, {Row: 0, Column: 1}, {Row: 0, Column: 2}, {Row: 0, Column: 3}, {Row: 0, Column: 4},
				{Row: 1, Column: 0}, {Row: 1, Column: 1}, {Row: 1, Column: 2}, {Row: 1, Column: 3}, {Row: 1, Column: 4},
				{Row: 2, Column: 0}, {Row: 2, Column: 1}, {Row: 2, Column: 2}, {Row: 2, Column: 3}, {Row: 2, Column: 4},
				{Row: 3, Column: 0}, {Row: 3, Column: 1}, {Row: 3, Column: 2}, {Row: 3, Column: 3}, {Row: 3, Column: 4},
			},
		}, {
			name: "test 4",
			args: args{
				bookedSeat: []*schema.Seat{{Row: 0, Column: 0}},
				Room:       &schema.Room{Row: 4, Column: 5, MinDistance: 6},
			},
			want: []*schema.Seat{{Row: 3, Column: 4}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetAvailableSeat(tt.args.bookedSeat, tt.args.Room); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAvailableSeat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkGetAvailableSeat(b *testing.B) {
	bookedSeats := []*schema.Seat{{Row: 1, Column: 1}, {Row: 1, Column: 2}}
	room := &schema.Room{Row: 10, Column: 10, MinDistance: 3}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		GetAvailableSeat(bookedSeats, room)
	}
}

func TestCalculateDistanceTwoSeat(t *testing.T) {
	type args struct {
		seat1 schema.Seat
		seat2 schema.Seat
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test 1",
			args: args{
				seat1: schema.Seat{Row: 1, Column: 1},
				seat2: schema.Seat{Row: 1, Column: 2},
			},
			want: 1,
		},
		{
			name: "test 2",
			args: args{
				seat1: schema.Seat{Row: 0, Column: 0},
				seat2: schema.Seat{Row: 4, Column: 4},
			},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalculateDistanceTwoSeat(tt.args.seat1, tt.args.seat2); got != tt.want {
				t.Errorf("CalculateDistanceTwoSeat() = %v, want %v", got, tt.want)
			}
		})
	}
}
