package app

import (
	"testing"

	"github.com/dattranman/simple_cinema/model/schema"
	"github.com/stretchr/testify/assert"
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
		want schema.SeatList
	}{
		{
			name: "test 3",
			args: args{
				bookedSeat: []*schema.Seat{},
				Room:       &schema.Room{Row: 4, Column: 5, MinDistance: 6},
			},
			want: schema.SeatList{
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
			want: schema.SeatList{
				{Row: 3, Column: 4},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetAvailableSeat(tt.args.bookedSeat, tt.args.Room)
			assert.Equal(t, tt.want, got)
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
