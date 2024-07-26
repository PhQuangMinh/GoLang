package rabbitmq

import (
	"fmt"
	"github.com/goccy/go-json"
	"github.com/streadway/amqp"
	"time"
)

func GetQueue(nameQueue string) {
	fmt.Println("Go RabbitMQ")

	conn, err := amqp.Dial("amqps://dgqdeyun:JQ3bkX-hrfUV0CD8FTMq_Zdtry-eijP3@armadillo.rmq.cloudamqp.com/dgqdeyun")
	if err != nil {
		fmt.Println("Failed to connect to RabbitMQ")
	}
	defer conn.Close()
	fmt.Println("Successfully")

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
			fmt.Printf("%s\n", d.Body)
			body, err := json.Marshal(d.Body)
			if err != nil {
				fmt.Println(err)
			}
			time.Sleep(3 * time.Second)
			fmt.Println("Day")
			Push(body, "QueueResult")
		}
	}()

	fmt.Println("Thanh cong")
}
