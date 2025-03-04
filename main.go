package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/dattranman/simple_cinema/api"
	"github.com/dattranman/simple_cinema/app"
	"github.com/dattranman/simple_cinema/config"
	"github.com/dattranman/simple_cinema/util"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/urfave/cli"
)

func runServer(c *cli.Context) error {
	configFile := c.String("config")
	cfg, err := config.Load(configFile)
	if err != nil {
		return err
	}
	app, err := app.New(cfg)
	if err != nil {
		return err
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(cors.New(cors.Config{
		AllowOrigins:              []string{"*"},
		AllowMethods:              []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete, http.MethodOptions},
		AllowCredentials:          true,
		AllowHeaders:              []string{"*"},
		OptionsResponseStatusCode: http.StatusOK,
	}))
	api := api.Init(app, router)
	err = api.Run()
	if err != nil {
		return err
	}

	return nil
}

// @title           Vulcan Cinema API
// @version         1.0
// @description     This is a API for Vulcan Cinema.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html
func main() {
	app := cli.NewApp()
	app.Name = "st"
	app.Flags = []cli.Flag{
		util.StringFlag("CONFIG", "config", "configuration file path", "config/config.yaml"),
	}
	app.Action = runServer

	if err := app.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
