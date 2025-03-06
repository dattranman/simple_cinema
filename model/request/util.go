package request

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

type ID struct {
	Version string `uri:"version" binding:"required" example:"v1"`
	IDStr   string `uri:"id" binding:"required" example:"1"`
	ID      int    `uri:"-"`
}

func (r *ID) Bind(c *gin.Context) (err error) {
	err = c.ShouldBindUri(r)
	if err != nil {
		return err
	}
	r.ID, err = strconv.Atoi(r.IDStr)
	if err != nil {
		return err
	}
	return nil
}
