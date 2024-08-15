package callservice

import (
	models "PhoneCall/model"
	"PhoneCall/repository"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type CallService struct {
	CallRepo repository.CallRepo
}

func NewCallService(CallRepo repository.CallRepo) *CallService {
	return &CallService{CallRepo: CallRepo}
}

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

func (getService *CallService) UpdateCall() func(c *gin.Context) {
	return func(c *gin.Context) {
		var data *models.Call
		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error1": err.Error(),
			})
			return
		}
		callResult, err := getService.CallRepo.UpdateCall(c, data)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
		}
		c.JSON(http.StatusOK, callResult)
	}
}

func (getService *CallService) DeleteCall() func(c *gin.Context) {
	return func(c *gin.Context) {
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error1": err.Error(),
			})
			return
		}
		err = getService.CallRepo.DeleteCall(id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
		}
		c.JSON(http.StatusOK, "Deleted successfully")
	}
}
