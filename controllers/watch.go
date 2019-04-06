package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
	"watchmen/model"
	"watchmen/services"
	"watchmen/services/watcher"
	"watchmen/services/watcher/watcherOperation"
)

func getDelay(c *gin.Context) time.Duration {
	maxDelay, hasDelay := c.GetQuery("max_delay")
	if !hasDelay {
		return 0
	}
	num, err := strconv.ParseInt(maxDelay, 10, 32)
	if err != nil {
		return 0
	}
	return time.Duration(num)
}

type WatchController struct{}

func doRequest(userData model.UserData, operation watcherOperation.Operation) error {
	var err error = nil
	if operation == watcherOperation.CheckIn {
		_, err = watcher.CheckIn(userData.User, userData.Password, watcher.WithCompany(userData.Company))
	} else {
		_, err = watcher.CheckOut(userData.User, userData.Password, watcher.WithCompany(userData.Company))
	}
	if err != nil {
		services.Logger.Error(err.Error(), userData.User, userData.Company)
		return err
	}
	response := gin.H{"status": "complete", "operation": "checkin", "company": userData.Company, "user": userData.User}
	services.Logger.Info(response)
	return nil
}

func doOperation(operation watcherOperation.Operation, c *gin.Context) {
	var userData model.UserData

	err := c.BindJSON(&userData)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		services.Logger.Error(err.Error(), userData.User, userData.Company)
		return
	}
	delay := getDelay(c)
	var response = gin.H{}
	if delay == 0 {
		err := doRequest(userData, operation)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		response = gin.H{"status": "complete", "operation": operation, "company": userData.Company, "user": userData.User}
	} else {
		response = gin.H{"status": "pending", "delay": delay, "operation": operation, "company": userData.Company, "user": userData.User}
		time.AfterFunc(delay*time.Minute, func() {
			_ = doRequest(userData, operation)
		})
	}
	services.Logger.Info(response)
	c.JSON(http.StatusOK, response)
}

// WatchController godoc
// @Summary Perform checkin via timewatch.co.il
// @Description get request to perform checkin
// @ID checkin
// @Produce json
// @Accept  json
// @Tags Watcher
// @Param user_data body model.UserData true "User Data"
// @Param max_delay query string false "Max delay in minutes until request is sent (random)"
// @Router /v1/watcher/checkin [post]
func (h WatchController) CheckIn(c *gin.Context) {
	doOperation(watcherOperation.CheckIn, c)
}

// WatchController godoc
// @Summary Perform checkout via timewatch.co.il
// @Description get request to perform checkout
// @ID checkout
// @Tags Watcher
// @Produce json
// @Accept json
// @Param user_data body model.UserData true "User Data"
// @Param max_delay query string false "Max delay in minutes until request is sent (random)"
// @Router /v1/watcher/checkout [post]
func (h WatchController) CheckOut(c *gin.Context) {
	doOperation(watcherOperation.CheckOut, c)
}
