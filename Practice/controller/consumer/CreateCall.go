package consumer

import (
	"Practice/controller"
	"Practice/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"github.com/streadway/amqp"
	"gorm.io/gorm"
	"net/http"
)

func CreateNewCall(nameQueue string) {
	var db = controller.MakeGorm()

	ch, con := controller.MakeChannel(nameQueue)
	defer ch.Close()
	defer con.Close()

	r := gin.Default()
	r.POST("/v1/items/", Create(db, ch, nameQueue))
	r.Run()
}

func Create(db *gorm.DB, ch *amqp.Channel, nameQueue string) func(context *gin.Context) {
	return func(c *gin.Context) {
		var data model.Call

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

		body, err := json.Marshal(data)

		err = ch.Publish(
			"",
			nameQueue,
			false,
			false,
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        body,
			},
		)
		if err != nil {
			fmt.Println(err)
			panic(err)
		}
		fmt.Println("Sent Message")
	}
}
