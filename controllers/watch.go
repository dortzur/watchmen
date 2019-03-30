package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"watchmen/model"
	"watchmen/services/watcher"
)

type WatchController struct{}

// WatchController godoc
// @Summary Perform checkin via timewatch.co.il
// @Description get request to perform checkin
// @ID checkin
// @Produce json
// @Accept  json
// @Tags Watcher
// @Param user_data body model.UserData true "User Data"
// @Router /v1/watcher/checkin [post]
func (h WatchController) CheckIn(c *gin.Context) {
	var userData model.UserData

	err := c.BindJSON(&userData)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err = watcher.CheckIn(userData.User, userData.Password, watcher.WithCompany(userData.Company))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "operation": "checkin", "company": userData.Company, "user": userData.User})
}

// WatchController godoc
// @Summary Perform checkout via timewatch.co.il
// @Description get request to perform checkout
// @ID checkout
// @Tags Watcher
// @Produce json
// @Accept json
// @Param user_data body model.UserData true "User Data"
// @Router /v1/watcher/checkout [post]
func (h WatchController) CheckOut(c *gin.Context) {
	company := c.Param("company")
	user := c.Param("username")
	pass := c.Param("password")
	_, err := watcher.CheckOut(user, pass, watcher.WithCompany(company))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "operation": "checkout", "company": company, "user": user})
}
