package callservice

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func (getService *GetCallService) GetCallsTime() func(ctx *gin.Context) {
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
