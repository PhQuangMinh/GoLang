package callservice

import (
	"PhoneCall/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

func (getService *CallService) GetCallsTime() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		layout := "2006-01-02"
		startAtStr := ctx.Query("startAt")
		startAt, err := time.Parse(layout, startAtStr)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		endAtStr := ctx.Query("endAt")
		endAt, err := time.Parse(layout, endAtStr)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		var paging model.Paging
		if err := ctx.ShouldBind(&paging); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		paging.Process()

		calls, err := getService.CallRepo.GetCalls(startAt, endAt, paging)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error3": err.Error(),
			})
		}
		ctx.JSON(http.StatusOK, calls)
	}
}

func (getService *CallService) GetValueField() func(c *gin.Context) {
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
