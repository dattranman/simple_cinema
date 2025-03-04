package api

import (
	"net/http"

	"github.com/dattranman/simple_cinema/model/request"
	"github.com/dattranman/simple_cinema/model/response"
	"github.com/gin-gonic/gin"
)

func (api *API) InitSeat() {
	api.RoomRouter.POST("/:id/seats", api.BookingSeat)
	api.RoomRouter.GET("/:id/available-seats", api.GetAvailableSeats)
	api.RoomRouter.DELETE("/:id/seats", api.CancelSeat)
}

// BookingSeat booking seat
// @Summary Booking seat
// @Description Booking seat
// @Accept json
// @Produce json
// @Param id path string true "Room ID"
// @Param seats body request.BookingSeats true "Seats"
// @Success 200 {object} response.BookingSeats
// @Failure 400 {object} response.Base
// @Failure 500 {object} response.Base
// @Router /api/v1/rooms/{id}/seats [post]
func (api *API) BookingSeat(c *gin.Context) {
	req := &request.BookingSeats{}
	err := req.Bind(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Base{
			Code:    response.CodeValidationError,
			Message: err.Error(),
		})
		return
	}

	resp, err := api.App.BookingSeat(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.Base{
			Code:    response.CodeError,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// GetAvailableSeats get available seats
// @Summary Get available seats
// @Description Get available seats
// @Accept json
// @Produce json
// @Param id path string true "Room ID"
// @Success 200 {object} response.GetAvailableSeats
// @Failure 400 {object} response.Base
// @Failure 500 {object} response.Base
// @Router /api/v1/rooms/{id}/available-seats [get]
func (api *API) GetAvailableSeats(c *gin.Context) {
	req := &request.GetAvailableSeats{}
	err := req.Bind(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Base{
			Code:    response.CodeValidationError,
			Message: err.Error(),
		})
		return
	}

	resp, err := api.App.GetAvailableSeats(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.Base{
			Code:    response.CodeError,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// CancelSeat cancel seat
// @Summary Cancel seat
// @Description Cancel seat
// @Param id path string true "Room ID"
// @Param seats body request.DeleteSeat true "Seats"
// @Success 200 {object} response.Base
// @Failure 400 {object} response.Base
// @Failure 500 {object} response.Base
// @Router /api/v1/rooms/{id}/seats [delete]
func (api *API) CancelSeat(c *gin.Context) {
	req := &request.DeleteSeat{}
	err := req.Bind(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Base{
			Code:    response.CodeValidationError,
			Message: err.Error(),
		})
		return
	}

	resp, err := api.App.CancelSeat(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.Base{
			Code:    response.CodeError,
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, resp)
}
