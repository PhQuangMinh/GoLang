package callservice

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

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
