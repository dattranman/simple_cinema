package request

import (
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
)

type BookingSeats struct {
	Version string `json:"-" uri:"version" binding:"required"`
	RoomID  int    `json:"-" uri:"id" binding:"required"`
	Seats   []Seat `json:"seats"`
}

type Seat struct {
	Row    uint `json:"row"`
	Column uint `json:"column"`
}

func (r *BookingSeats) Bind(c *gin.Context) error {
	err := c.ShouldBindUri(&r)
	if err != nil {
		return err
	}
	if err := c.ShouldBindJSON(&r); err != nil {
		return err
	}
	if len(r.Seats) == 0 {
		return errors.New("seats is required")
	}

	// Check for duplicate seats
	seenSeats := make(map[string]bool)
	for _, seat := range r.Seats {
		key := fmt.Sprintf("%d-%d", seat.Row, seat.Column)
		if seenSeats[key] {
			return errors.New("duplicate seats are not allowed")
		}
		seenSeats[key] = true
	}
	return nil
}

type GetAvailableSeats struct {
	Version string `form:"version" uri:"version" binding:"required"`
	RoomID  int    `form:"id" uri:"id" binding:"required"`
}

func (r *GetAvailableSeats) Bind(c *gin.Context) (err error) {
	err = c.ShouldBindUri(r)
	if err != nil {
		return err
	}
	return nil
}

type DeleteSeat struct {
	Version string `json:"-" uri:"version" binding:"required"`
	RoomID  int    `json:"-" uri:"id" binding:"required"`
	Seats   []Seat `json:"seats"`
}

func (r *DeleteSeat) Bind(c *gin.Context) error {
	err := c.ShouldBindUri(&r)
	if err != nil {
		return err
	}
	if err := c.ShouldBindJSON(&r); err != nil {
		return err
	}
	if len(r.Seats) == 0 {
		return errors.New("seats is required")
	}
	return nil
}
