package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"watchmen/services/watcher"
)

type WatchController struct{}

// WatchController godoc
// @Summary Perform Checkin to timewatch.co.il
// @Description get request to perform checkin
// @ID checkin
// @Produce json
// @Tags Watcher
// @Param company_id path string true "Company ID"
// @Param employee_id path string true "Employee ID"
// @Param password path string true "Employee Password"
// @Router /v1/watcher/checkin/{company_id}/{employee_id}/{password} [get]
func (h WatchController) CheckIn(c *gin.Context) {
	company := c.Param("company")
	user := c.Param("username")
	pass := c.Param("password")

	_, err := watcher.CheckIn(user, pass, watcher.WithCompany(company))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "operation": "checkin", "company": company, "user": user})
}

// WatchController godoc
// @Summary Perform Checkout to timewatch.co.il
// @GroupName baba
// @Description get request to perform checkout
// @ID checkout
// @Tags Watcher
// @Produce json
// @Param company_id path string true "Company ID"
// @Param employee_id path string true "Employee ID"
// @Param password path string true "Employee Password"
// @Router /v1/watcher/checkout/{company_id}/{employee_id}/{password} [get]
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
