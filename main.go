package main

import (
	"watchmen/config"
	"watchmen/docs"
	"watchmen/server"
	"watchmen/services"
)

// @version 1.0
func main() {
	config.Init(services.GetEnv())
	docs.SwaggerInfo.Title = "Watchmen timewatch.co.il API"
	docs.SwaggerInfo.Description = "https://github.com/dortzur/watchmen"
	server.Init()

}
