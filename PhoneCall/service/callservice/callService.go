package callservice

import (
	models "PhoneCall/model"
	"PhoneCall/repository"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type GetCallService struct {
	CallRepo repository.CallRepo
}

func NewGetCallService(CallRepo repository.CallRepo) *GetCallService {
	return &GetCallService{CallRepo: CallRepo}
}

func (getService *GetCallService) GetValueField() func(c *gin.Context) {
	return func(c *gin.Context) {
		idStr := c.Query("id")
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		field := c.Query("metadata_display_field")
		valueField, err := getService.CallRepo.GetValueField(id, field)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, valueField)
	}
}

func (getService *GetCallService) CreateNewCall() func(c *gin.Context) {
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

func (getService *GetCallService) UpdateCall() func(c *gin.Context) {
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

func (getService *GetCallService) DeleteCall() func(c *gin.Context) {
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
