package callservice

import (
	models "PhoneCall/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (getService *CallService) CreateNewCall() func(c *gin.Context) {
	return func(c *gin.Context) {
		var data *models.Call
		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error1": err.Error(),
			})
			return
		}
		callResult, err := getService.CallRepo.CreateNewCall(data)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
		}
		c.JSON(http.StatusOK, callResult)
	}
}
