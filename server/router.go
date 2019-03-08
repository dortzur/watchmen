package server

import (
	"github.com/gin-gonic/gin"
	"watchmen/controllers"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	health := new(controllers.HealthController)
	router.GET("/health", health.Status)

	v1 := router.Group("v1")
	{
		watcherGroup := v1.Group("watcher")
		{
			watcher := new(controllers.WatchController)
			watcherGroup.GET("checkin/:username/:password", watcher.CheckIn)
			watcherGroup.GET("checkout/:username/:password", watcher.CheckIn)
		}
	}

	return router
}
