package repository

import (
	"Practice/model"
	"github.com/streadway/amqp"
)

type RabbitMQRepo interface {
	Push(nameQueue string, call model.Call)
	Pop(nameQueue string) (<-chan amqp.Delivery, error)
}
