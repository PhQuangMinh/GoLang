package repoimpl

import (
	models "PhoneCall/models"
	"encoding/json"
	"fmt"
	"github.com/streadway/amqp"
)

type RabbitMQ struct {
	Channel *amqp.Channel
	Connect *amqp.Connection
}

func NewRabbitMQ(channel *amqp.Channel, connect *amqp.Connection) *RabbitMQ {
	return &RabbitMQ{
		Channel: channel,
		Connect: connect,
	}
}

func (rabbit *RabbitMQ) Push(nameQueue string, call models.Call) {
	body, err := json.Marshal(call)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = rabbit.Channel.Publish(
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
		return
	}
}

func (rabbit *RabbitMQ) Pop(nameQueue string) (<-chan amqp.Delivery, error) {
	messages, err := rabbit.Channel.Consume(
		nameQueue,
		"",
		true,
		false,
		false,
		false,
		nil)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return messages, nil
}
