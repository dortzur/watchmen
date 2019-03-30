package server

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"net/http"
	"watchmen/controllers"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	health := new(controllers.HealthController)
	router.GET("/health", health.Status)
	router.GET("/", func(context *gin.Context) {
		context.Redirect(http.StatusPermanentRedirect, "/swagger/index.html")
	})
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	v1 := router.Group("v1")
	{
		watcherGroup := v1.Group("watcher")
		{
			watcher := new(controllers.WatchController)
			watcherGroup.POST("checkin", watcher.CheckIn)
			watcherGroup.POST("checkout", watcher.CheckOut)
		}
	}
	return router
}
