package consumercreate

import (
	"Practice/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"github.com/streadway/amqp"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/http"
)

func makeGorm() *gorm.DB {
	dsn := "root:123456@tcp(127.0.0.1:3306)/call_management?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Failed connect database")
		return nil
	}
	return db
}

func makeChannel(nameQueue string) (*amqp.Channel, *amqp.Connection) {
	conn, err := amqp.Dial("amqps://dgqdeyun:JQ3bkX-hrfUV0CD8FTMq_Zdtry-eijP3@armadillo.rmq.cloudamqp.com/dgqdeyun")
	if err != nil {
		fmt.Println(err)
	}

	ch, err := conn.Channel()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	q, err := ch.QueueDeclare(
		nameQueue,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	fmt.Println(q)
	return ch, conn
}
func CreateNewCall(nameQueue string) {
	var db = makeGorm()

	ch, con := makeChannel(nameQueue)
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
