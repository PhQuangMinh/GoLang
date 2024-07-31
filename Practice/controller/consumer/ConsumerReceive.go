package consumer

import (
	"Practice/controller"
	"Practice/model"
	"fmt"
	"github.com/goccy/go-json"
	"github.com/streadway/amqp"
	"time"
)

func receiveQueueSolve(ch *amqp.Channel, nameQueue string, waitTime time.Duration) {
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
			cal.CallAnsweredTime = time.Now()
			time.Sleep(waitTime)
			cal.CallEndedTime = time.Now()
			if er != nil {
				fmt.Println(er)
			}
			fmt.Println(cal)
			pushQueueResult(ch, "QueueResult", cal)
			cha <- true
		}
	}()
	<-cha
	fmt.Println("Thanh cong")
}

func pushQueueResult(ch *amqp.Channel, nameQueue string, data interface{}) {
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
	fmt.Println("Sent Message Queue Result")
}

func ReceiveCall(nameQueue string) {
	ch, con := controller.MakeChannel(nameQueue)
	defer ch.Close()
	defer con.Close()

	receiveQueueSolve(ch, nameQueue, 5)
}
