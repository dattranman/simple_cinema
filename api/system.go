package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (api *API) HealthCheck() {
	api.BaseRouters.Root.GET("", api.healthCheck)
}

// @BasePath /api/:version

// HealthCheck godoc
// @Summary check health
// @Schemes
// @Description check health
// @Accept json
// @Produce json
// @Success 200 {string} start at port :port
// @Router / [get]
func (api *API) healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, "start at port "+api.App.Config.ServiceSettings.Port)

}
