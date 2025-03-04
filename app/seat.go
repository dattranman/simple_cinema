package app

import (
	"fmt"
	"math"

	"github.com/dattranman/simple_cinema/model/request"
	"github.com/dattranman/simple_cinema/model/response"
	"github.com/dattranman/simple_cinema/model/schema"
)

type Seat struct {
	Row    int `json:"row"`
	Column int `json:"column"`
}

func CalculateDistanceTwoSeat(seat1, seat2 schema.Seat) int {
	result := int(math.Abs(float64(seat1.Row-seat2.Row)) + math.Abs(float64(seat1.Column-seat2.Column)))
	return result
}

func CalculateMinDistanceFromTwoGroupSeat(group1, group2 []*schema.Seat) int {
	minDistance := math.MaxInt
	for _, seat1 := range group1 {
		for _, seat2 := range group2 {
			distance := CalculateDistanceTwoSeat(*seat1, *seat2)
			if distance < minDistance {
				minDistance = distance
			}
		}
	}
	return minDistance
}

func GetAvailableSeat(bookedSeat []*schema.Seat, room *schema.Room) []*schema.Seat {
	availableSeat := []*schema.Seat{}
	for row := 0; row < room.Row; row++ {
		for col := 0; col < room.Column; col++ {
			seat := schema.Seat{Row: row, Column: col}
			if !contains(bookedSeat, seat) && isSafeForBooking(seat, bookedSeat, room.MinDistance) {
				availableSeat = append(availableSeat, &seat)
			}
		}
	}
	return availableSeat
}

func GetAvailableSeatOptimized(bookedSeat []*schema.Seat, room *schema.Room) (availableSeats []*schema.Seat) {
	const formatKeySeat = "%d-%d"
	mapSeatCannotBeUsed := make(map[string]bool)
	for _, seat := range bookedSeat {
		mapSeatCannotBeUsed[fmt.Sprintf(formatKeySeat, seat.Row, seat.Column)] = true
		for i := 1; i <= room.MinDistance; i++ {
			if seat.Row+i < room.Row {
				mapSeatCannotBeUsed[fmt.Sprintf(formatKeySeat, seat.Row+i, seat.Column)] = true
			}
			if seat.Row-i >= 0 {
				mapSeatCannotBeUsed[fmt.Sprintf(formatKeySeat, seat.Row-i, seat.Column)] = true
			}
			if seat.Column+i < room.Column {
				mapSeatCannotBeUsed[fmt.Sprintf(formatKeySeat, seat.Row, seat.Column+i)] = true
			}
			if seat.Column-i >= 0 {
				mapSeatCannotBeUsed[fmt.Sprintf(formatKeySeat, seat.Row, seat.Column-i)] = true
			}
		}
	}
	for row := 0; row < room.Row; row++ {
		for col := 0; col < room.Column; col++ {
			if mapSeatCannotBeUsed[fmt.Sprintf(formatKeySeat, row, col)] {
				continue
			}
			availableSeats = append(availableSeats, &schema.Seat{Row: row, Column: col})
		}
	}
	return availableSeats
}

func isSafeForBooking(seat schema.Seat, bookedSeats []*schema.Seat, minDistance int) bool {
	for _, bookedSeat := range bookedSeats {
		distance := CalculateDistanceTwoSeat(seat, *bookedSeat)
		if distance > minDistance {
			return true
		}
	}
	return false
}

func contains(seats []*schema.Seat, seat schema.Seat) bool {
	for _, s := range seats {
		if s.Row == seat.Row && s.Column == seat.Column {
			return true
		}
	}
	return false
}

func (app *App) BookingSeat(req *request.BookingSeats) (resp *response.BookingSeats, err error) {
	room, err := app.Store.Room().GetByID(req.RoomID)
	if err != nil {
		return nil, err
	}
	bookedSeat, err := app.Store.Seat().GetByRoomID(room.ID)
	if err != nil {
		return nil, err
	}
	for _, seat := range req.Seats {
		if int(seat.Row) >= room.Row || int(seat.Column) >= room.Column {
			resp = &response.BookingSeats{
				Base: response.Base{
					Message: "Seat is out of range",
					Code:    response.CodeError,
				},
			}
			return resp, nil
		}
		if contains(bookedSeat, schema.Seat{Row: int(seat.Row), Column: int(seat.Column)}) {
			resp = &response.BookingSeats{
				Base: response.Base{
					Message: "Seat is already booked",
					Code:    response.CodeError,
				},
			}
			return resp, nil
		}
		if !isSafeForBooking(schema.Seat{Row: int(seat.Row), Column: int(seat.Column)}, bookedSeat, room.MinDistance) {
			resp = &response.BookingSeats{
				Base: response.Base{
					Message: "Seat is not safe for booking",
					Code:    response.CodeError,
				},
			}
			return resp, nil
		}
	}
	dataSeat := make([]*schema.Seat, len(req.Seats))
	for i, seat := range req.Seats {
		dataSeat[i] = &schema.Seat{
			RoomID: room.ID,
			Row:    int(seat.Row),
			Column: int(seat.Column),
		}
	}
	err = app.Store.Seat().Create(dataSeat)
	if err != nil {
		return nil, err
	}
	responseSeats := make([]response.Seat, len(req.Seats))
	for i, seat := range dataSeat {
		responseSeats[i] = seat.ToResponse()
	}
	return &response.BookingSeats{
		Base: response.Base{
			Message: "Seat booked successfully",
			Code:    response.CodeSuccess,
		},
		Seats: responseSeats,
	}, nil
}

func (app *App) GetAvailableSeats(req *request.GetAvailableSeats) (resp *response.GetAvailableSeats, err error) {
	room, err := app.Store.Room().GetByID(req.RoomID)
	if err != nil {
		return nil, err
	}
	bookedSeat, err := app.Store.Seat().GetByRoomID(room.ID)
	if err != nil {
		return nil, err
	}
	availableSeat := GetAvailableSeat(bookedSeat, room)
	responseAvailableSeat := make([]response.Seat, len(availableSeat))
	for i, seat := range availableSeat {
		responseAvailableSeat[i] = seat.ToResponse()
	}
	return &response.GetAvailableSeats{
		Base: response.Base{
			Message: "Available seats",
			Code:    response.CodeSuccess,
		},
		AvailableSeats: responseAvailableSeat,
	}, nil
}

func (app *App) CancelSeat(req *request.DeleteSeat) (resp *response.Base, err error) {
	room, err := app.Store.Room().GetByID(req.RoomID)
	if err != nil {
		return nil, err
	}
	dataSeat := make([]*schema.Seat, len(req.Seats))
	for i, seat := range req.Seats {
		dataSeat[i] = &schema.Seat{
			RoomID: room.ID,
			Row:    int(seat.Row),
			Column: int(seat.Column),
		}
	}
	err = app.Store.Seat().Delete(room.ID, dataSeat)
	if err != nil {
		return nil, err
	}
	return &response.Base{
		Message: "Seat canceled successfully",
		Code:    response.CodeSuccess,
	}, nil
}
