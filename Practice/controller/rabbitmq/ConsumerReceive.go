package rabbitmq

import (
	"Practice/model"
	"fmt"
	"github.com/goccy/go-json"
	"github.com/streadway/amqp"
	"time"
)

func GetQueue(nameQueue string, waitTime time.Duration) {
	conn, err := amqp.Dial("amqps://dgqdeyun:JQ3bkX-hrfUV0CD8FTMq_Zdtry-eijP3@armadillo.rmq.cloudamqp.com/dgqdeyun")
	if err != nil {
		fmt.Println("Failed to connect to RabbitMQ")
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer ch.Close()

	msgs, err := ch.Consume(
		nameQueue,
		"",
		true,
		false,
		false,
		false,
		nil)

	go func() {
		for d := range msgs {
			var cal model.Call
			er := json.Unmarshal(d.Body, &cal)
			cal.CallAnsweredTime = time.Now()
			time.Sleep(waitTime)
			cal.CallEndedTime = time.Now()
			fmt.Println(er)
			if err != nil {
				fmt.Println(err)
			}
			Push(cal, "QueueResult")
		}
	}()

	fmt.Println("Thanh cong")
}

func GetQueueResult(nameQueue string) {
	conn, err := amqp.Dial("amqps://dgqdeyun:JQ3bkX-hrfUV0CD8FTMq_Zdtry-eijP3@armadillo.rmq.cloudamqp.com/dgqdeyun")
	if err != nil {
		fmt.Println("Failed to connect to RabbitMQ")
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer ch.Close()

	msgs, err := ch.Consume(
		nameQueue,
		"",
		true,
		false,
		false,
		false,
		nil)

	fmt.Println("Queue Result")
	go func() {
		for d := range msgs {
			var cal model.Call
			er := json.Unmarshal(d.Body, &cal)
			fmt.Println(er, cal)
		}
	}()

}
