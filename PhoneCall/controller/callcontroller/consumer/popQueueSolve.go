package consumer

import (
	"PhoneCall/common"
	"PhoneCall/repository"
	"PhoneCall/service/connection"
	"fmt"
	"github.com/goccy/go-json"
	"time"
)

func popQueueSolve(rabbit *repository.RabbitMQ, nameQueue string, waitTime time.Duration) {
	messages, err := rabbit.Pop(nameQueue)
	if err != nil {
		fmt.Println(err)
		return
	}
	cha := make(chan bool)

	go func() {
		for d := range messages {
			var cal modelss.Call
			er := json.Unmarshal(d.Body, &cal)

			cal.CallAnsweredTime = time.Now()
			time.Sleep(waitTime)
			cal.CallEndedTime = time.Now()

			if er != nil {
				fmt.Println(er)
				continue
			}
			rabbit.Push("QueueResult", cal)
			cha <- true
		}
	}()
	<-cha
	fmt.Println("Thanh cong")
}

func ReceiveCall(nameQueue string) {
	channel, connect, err := connection.ConnectRabbit("QueueSolve", common.LinkRabbit)
	if err != nil {
		fmt.Println(err)
		return
	}
	rabbit := repository.NewRabbitMQ(channel, connect)
	defer rabbit.Connect.Close()
	defer rabbit.Channel.Close()

	popQueueSolve(rabbit, nameQueue, 5)
}
