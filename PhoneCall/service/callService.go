package service

import (
	"PhoneCall/repository"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

type GetCallService struct {
	CallRepo repository.CallRepo
}

func NewGetCallService(CallRepo repository.CallRepo) *GetCallService {
	return &GetCallService{CallRepo: CallRepo}
}

func (getService *GetCallService) GetCalls() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		layout := "2006-01-02"
		startAtStr := ctx.Query("startAt")
		startAt, err := time.Parse(layout, startAtStr)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error1": err.Error(),
			})
			return
		}

		endAtStr := ctx.Query("endAt")
		endAt, err := time.Parse(layout, endAtStr)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error2": err.Error(),
			})
			return
		}

		calls, err := getService.CallRepo.GetCalls(startAt, endAt)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error3": err.Error(),
			})
		}
		ctx.JSON(http.StatusOK, calls)
	}
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
