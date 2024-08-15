package connection

import (
	"github.com/streadway/amqp"
)

func ConnectRabbit(nameQueue, pathMQ string) (*amqp.Channel, *amqp.Connection, error) {
	conn, err := amqp.Dial(pathMQ)
	if err != nil {
		return nil, nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, nil, err
	}

	_, err = ch.QueueDeclare(
		nameQueue,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, nil, err
	}

	return ch, conn, nil
}
