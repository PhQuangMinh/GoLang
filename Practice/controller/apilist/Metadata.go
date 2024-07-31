package apilist

import (
	"Practice/controller"
	"Practice/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"strings"
)

func GetItem() {
	r := gin.Default()
	r.GET("/v1/items/:id", getItemByID(controller.MakeGorm()))
	r.Run()
}

func getItemByID(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		var data model.Call

		if err := db.Table("calls").Where("id = ?", id).First(&data).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		var result map[string]interface{}
		data_json, err := json.Marshal(data)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}
		if err := json.Unmarshal(data_json, &result); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}
		display_field := c.Query("metadata_display_field")

		if display_field == "" {
			c.JSON(http.StatusOK, gin.H{
				"data": result["metadata"],
			})
		} else {
			jsonString, ok := result["metadata"].(string)
			fmt.Println(ok, jsonString)
			index := strings.Index(jsonString, display_field)

			if index == -1 {
				c.JSON(http.StatusOK, gin.H{
					"data": "",
				})
			} else {
				value := ""
				for i := index + len(display_field) + 1; i < len(jsonString); i++ {
					if jsonString[i] == ',' || jsonString[i] == '}' {
						break
					}
					value = value + jsonString[i:i+1]
				}
				c.JSON(http.StatusOK, gin.H{
					"data": "{\"" + display_field + "\"" + value + "}",
				})
			}
		}
	}
}
