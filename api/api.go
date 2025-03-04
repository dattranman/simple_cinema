package api

import (
	"github.com/dattranman/simple_cinema/app"
	"github.com/dattranman/simple_cinema/docs"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Routers struct {
	Root *gin.Engine

	Api *gin.RouterGroup
}

type API struct {
	App         *app.App
	BaseRouters *Routers

	RoomRouter *gin.RouterGroup
}

func Init(app *app.App, router *gin.Engine) *API {
	api := &API{
		App: app,
		BaseRouters: &Routers{
			Root: router,
		},
	}

	docs.SwaggerInfo.BasePath = "/"
	api.BaseRouters.Root.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	api.BaseRouters.Api = api.BaseRouters.Root.Group("/api/:version")
	api.RoomRouter = api.BaseRouters.Api.Group("/rooms")
	api.HealthCheck()
	api.InitRoom()
	api.InitSeat()
	return api
}

func (a *API) Run() error {
	err := a.BaseRouters.Root.Run(a.App.Config.ServiceSettings.Port)
	if err != nil {
		return err
	}
	return nil
}
