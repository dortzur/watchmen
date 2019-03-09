package server

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"watchmen/controllers"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	health := new(controllers.HealthController)
	router.GET("/health", health.Status)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	v1 := router.Group("v1")
	{
		watcherGroup := v1.Group("watcher")
		{
			watcher := new(controllers.WatchController)
			watcherGroup.GET("checkin/:company/:username/:password", watcher.CheckIn)
			watcherGroup.GET("checkout/:company/:username/:password", watcher.CheckOut)
		}
	}
	return router
}
