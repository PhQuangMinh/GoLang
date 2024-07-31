package controller

import (
	"fmt"
	"github.com/streadway/amqp"
)

func MakeChannel(nameQueue string) (*amqp.Channel, *amqp.Connection) {
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
