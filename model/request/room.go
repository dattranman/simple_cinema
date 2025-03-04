package request

import (
	"github.com/gin-gonic/gin"
)

type CreateRoom struct {
	Row         uint `json:"row" binding:"required"`
	Column      uint `json:"column" binding:"required"`
	MinDistance uint `json:"min_distance"`
}

func (r *CreateRoom) Bind(c *gin.Context) error {
	if err := c.ShouldBindJSON(r); err != nil {
		return err
	}
	return nil
}

type UpdateRoom struct {
	ID     string `uri:"id" binding:"required"`
	Row    uint   `json:"row"`
	Column uint   `json:"column"`
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
	ID int `uri:"id" binding:"required"`
}

func (r *GetRoomDetail) Bind(c *gin.Context) error {
	if err := c.ShouldBindUri(r); err != nil {
		return err
	}
	return nil
}
