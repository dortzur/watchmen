package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type HealthController struct{}

func (h HealthController) Status(c *gin.Context) {
	res := gin.H{"status": "strong like a bull!"}
	c.JSON(http.StatusOK, res)
}
