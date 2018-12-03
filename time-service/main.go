package main

import (
	"net/http"
	"go-micro-demo1/time-service/API/config"
	"go-micro-demo1/time-service/API/router"
	"github.com/gin-gonic/gin"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"go-micro-demo1/time-service/server"
)

var (
	cfg = pflag.StringP("config", "c", "", "apiserver config file path.")
)
func main() {
	pflag.Parse()

	// init config
	if err := config.Init(*cfg); err != nil {
		panic(err)
	}

	// Set gin mode.
	gin.SetMode(viper.GetString("runmode"))

	// Create the Gin engine.
	g := gin.New()

	middlewares := []gin.HandlerFunc{}
	// Routes.
	router.Load(
		// Cores.
		g,
		// Middlwares.
		middlewares...,
	)
	go server.Server()
	http.ListenAndServe(viper.GetString("addr"), g)

}
