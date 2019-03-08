package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"watchmen/services/watcher"
)

type WatchController struct{}

func (h WatchController) CheckIn(c *gin.Context) {
	company := c.Query("company")
	user := c.Param("username")
	pass := c.Param("password")
	res, err := watcher.CheckIn(user, pass, watcher.WithCompany(company))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h WatchController) CheckOut(c *gin.Context) {
	company := c.Query("company")
	user := c.Param("username")
	pass := c.Param("password")
	res := gin.H{"result": watcher.CheckOut(user, pass, watcher.WithCompany(company))}
	c.JSON(http.StatusOK, res)
}
