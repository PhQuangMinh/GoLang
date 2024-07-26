package database

import (
	"Practice/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/http"
)

func PostData(data model.Call) {
	dsn := "root:123456@tcp(127.0.0.1:3306)/call_management?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Failed connect database")
	}

	r := gin.Default()

	r.POST("/v1/items/", CreateItem(db, data))
	r.Run()
}

func CreateItem(db *gorm.DB, data model.Call) func(context *gin.Context) {
	return func(c *gin.Context) {

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"Error": err.Error(),
			})
			return
		}

		if err := db.Create(&data).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, data)
	}
}
