package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type HealthController struct{}

// Health godoc
// @Router /health [get]
// @Summary Return Health Status
// @Description get health json
// @ID get-health
// @Accept  json
// @Produce  json
func (h HealthController) Status(c *gin.Context) {
	res := gin.H{"status": "strong like a bull!"}
	c.JSON(http.StatusOK, res)
}
