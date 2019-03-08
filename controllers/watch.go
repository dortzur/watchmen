package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"watchmen/services/watcher"
)

type WatchController struct{}

func (h WatchController) CheckIn(c *gin.Context) {
	defaultOptions := watcher.CheckIn("baba", "ganush")
	companyOptions := watcher.CheckIn("yo", "hey", watcher.WithCompany("eeee"))
	res := gin.H{"default": defaultOptions,
		"company": companyOptions,
	}
	c.JSON(http.StatusOK, res)
}
