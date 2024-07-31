package consumer

import (
	"Practice/controller"
	"Practice/model"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/streadway/amqp"
	"gorm.io/gorm"
	"net/http"
)

func receiveQueueResult(ch *amqp.Channel, nameQueue string, db *gorm.DB) {
	msgs, err := ch.Consume(
		nameQueue,
		"",
		true,
		false,
		false,
		false,
		nil)
	fmt.Println(err)
	cha := make(chan bool)

	go func() {
		for d := range msgs {
			var cal model.Call
			er := json.Unmarshal(d.Body, &cal)
			if er != nil {
				fmt.Println(er)
				continue
			}
			cal.CallResult = "Success"
			fmt.Println(cal)
			updateDatabase(db, cal)
		}
		cha <- true
	}()
	<-cha
	fmt.Println("Thanh cong")
}

func updateDatabase(db *gorm.DB, data model.Call) {
	r := gin.Default()
	r.PUT("/v1/items", func(c *gin.Context) {

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"Error": err.Error(),
			})
			return
		}

		if err := db.Table("calls").Where("id = ?", data.Id).Updates(&data).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data": data,
		})
	})
	r.Run()
}

func UpdateQueueResult(nameQueue string) {
	ch, conn := controller.MakeChannel(nameQueue)
	db := controller.MakeGorm()
	defer ch.Close()
	defer conn.Close()

	receiveQueueResult(ch, nameQueue, db)
}
