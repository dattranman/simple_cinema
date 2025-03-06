package request

import (
	"github.com/gin-gonic/gin"
)

type CreateRoom struct {
	Row         uint `json:"row" binding:"required" example:"4"`
	Column      uint `json:"column" binding:"required" example:"5"`
	MinDistance uint `json:"min_distance" example:"6"`
}

func (r *CreateRoom) Bind(c *gin.Context) error {
	if err := c.ShouldBindJSON(r); err != nil {
		return err
	}
	return nil
}

type UpdateRoom struct {
	ID          string `uri:"id" binding:"required" example:"1"`
	Row         uint   `json:"row" example:"4"`
	Column      uint   `json:"column" example:"5"`
	MinDistance uint   `json:"min_distance" example:"6"`
}

func (r *UpdateRoom) Bind(c *gin.Context) error {
	if err := c.ShouldBindUri(r); err != nil {
		return err
	}
	if err := c.ShouldBindJSON(r); err != nil {
		return err
	}
	return nil
}

type GetRoomDetail struct {
	ID int `uri:"id" binding:"required" example:"1"`
}

func (r *GetRoomDetail) Bind(c *gin.Context) error {
	if err := c.ShouldBindUri(r); err != nil {
		return err
	}
	return nil
}

type DeleteRoom struct {
	ID int `uri:"id" binding:"required" example:"1"`
}

func (r *DeleteRoom) Bind(c *gin.Context) error {
	if err := c.ShouldBindUri(r); err != nil {
		return err
	}
	return nil
}
