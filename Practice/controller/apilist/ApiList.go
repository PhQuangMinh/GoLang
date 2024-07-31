package apilist

import (
	"Practice/controller"
	"Practice/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func GetList() {
	r := gin.Default()
	r.GET("/v1/items", listItem(controller.MakeGorm()))
	r.Run()
}

func listItem(db *gorm.DB) func(context *gin.Context) {
	return func(c *gin.Context) {

		var result []model.Call

		start := c.Query("start")
		end := c.Query("end")
		if err := db.Table("calls").
			Order("id, client_name, phone_number").
			Where("created_at between ? and ?", start, end).
			Find(&result).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data": result,
		})
	}
}
