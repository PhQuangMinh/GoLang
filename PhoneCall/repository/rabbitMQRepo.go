package repository

import (
	models "PhoneCall/models"
	"github.com/streadway/amqp"
)

type RabbitMQRepo interface {
	Push(nameQueue string, call models.Call)
	Pop(nameQueue string) (<-chan amqp.Delivery, error)
}
