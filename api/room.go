package api

import (
	"net/http"

	"github.com/dattranman/simple_cinema/model/request"
	"github.com/dattranman/simple_cinema/model/response"
	"github.com/gin-gonic/gin"
)

func (api *API) InitRoom() {
	api.RoomRouter.GET("", api.GetRooms)
	api.RoomRouter.GET("/:id", api.GetRoomDetail)
	api.RoomRouter.POST("", api.CreateRoom)
	api.RoomRouter.DELETE("/:id", api.DeleteRoom)
}

// GetRooms get all rooms
// @Summary Get all rooms
// @Description Get all rooms
// @Accept json
// @Produce json
// @Success 200 {object} response.Base
// @Failure 500 {object} response.Base
// @Router /api/v1/rooms [get]
func (api *API) GetRooms(c *gin.Context) {
	rooms, err := api.App.GetRoomList()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, rooms)
}

// GetRoomDetail get room detail
// @Summary Get room detail
// @Description Get room detail
// @Accept json
// @Produce json
// @Param id path int true "Room ID"
// @Success 200 {object} response.Base
// @Failure 400 {object} response.Base
// @Failure 500 {object} response.Base
// @Router /api/v1/rooms/{id} [get]
func (api *API) GetRoomDetail(c *gin.Context) {
	var req request.GetRoomDetail
	err := req.Bind(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Base{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}
	resp, err := api.App.GetRoomDetail(req.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.Base{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	c.JSON(http.StatusOK, resp)
}

// CreateRoom create room
// @Summary Create room
// @Description Create room
// @Accept json
// @Produce json
// @Param room body request.CreateRoom true "Room"
// @Success 200 {object} response.Base
// @Failure 400 {object} response.Base
// @Failure 500 {object} response.Base
// @Router /api/v1/rooms [post]
func (api *API) CreateRoom(c *gin.Context) {
	var req request.CreateRoom
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.Base{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	room, err := api.App.CreateRoom(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.Base{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	c.JSON(http.StatusOK, room)
}

// DeleteRoom delete room
// @Summary Delete room
// @Description Delete room
// @Accept json
// @Produce json
// @Param id path int true "Room ID"
// @Success 200 {object} response.Base
// @Failure 400 {object} response.Base
// @Failure 500 {object} response.Base
// @Router /api/v1/rooms/{id} [delete]
func (api *API) DeleteRoom(c *gin.Context) {
	var req request.DeleteRoom
	err := req.Bind(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Base{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	err = api.App.DeleteRoom(req.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.Base{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}
	c.JSON(http.StatusOK, response.Base{
		Code:    http.StatusOK,
		Message: "Room deleted successfully",
	})
}
