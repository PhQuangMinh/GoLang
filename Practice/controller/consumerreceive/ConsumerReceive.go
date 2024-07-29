package consumerreceive

import (
	"Practice/model"
	"fmt"
	"github.com/goccy/go-json"
	"github.com/streadway/amqp"
	"time"
)

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
	ch, con := makeChannel(nameQueue)
	defer ch.Close()
	defer con.Close()

	receiveQueueSolve(ch, nameQueue, 5)
}
