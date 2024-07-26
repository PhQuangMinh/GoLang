package main

import (
	"fmt"
	"github.com/streadway/amqp"
)

func main() {
	// 1. Thiết lập kết nối tới máy chủ RabbitMQ
	conn, err := amqp.Dial("amqps://dgqdeyun:JQ3bkX-hrfUV0CD8FTMq_Zdtry-eijP3@armadillo.rmq.cloudamqp.com/dgqdeyun")
	if err != nil {
		fmt.Println("Failed to connect to RabbitMQ")
	}
	defer conn.Close()
	fmt.Println("Successfully")

	// 2. Tạo một kênh
	ch, err := conn.Channel()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer ch.Close()

	// 3. Khai báo một hàng đợi
	q, err := ch.QueueDeclare(
		"Test queue", // Tên hàng đợi
		false,        // Durable - hàng đợi có bền bỉ không
		false,        // Delete when unused - xóa hàng đợi khi không sử dụng
		false,        // Exclusive - hàng đợi chỉ dành riêng cho kết nối này
		false,        // No-wait - không chờ cho đến khi hàng đợi được khai báo thành công
		nil,          // Arguments - các tham số bổ sung
	)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	fmt.Println(q)

	// 4. Gửi một thông điệp tới hàng đợi
	err = ch.Publish(
		"",           // Exchange - trao đổi (exchange) để gửi thông điệp, ở đây để trống
		"Test queue", // Routing key - khóa định tuyến, ở đây là tên hàng đợi
		false,        // Mandatory - thông điệp phải được xác đinh tới hàng đợi, nếu không thì trả về cho nhà sản xuất
		false,        // Immediate - thông điệp phải được tiêu thụ ngay lập tức hoặc bị trả lại cho nhà sản xuất
		amqp.Publishing{
			ContentType: "text/plain",          // Loại nội dung của thông điệp
			Body:        []byte("Hello World"), // Nội dung của thông điệp
		},
	)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	fmt.Println("Gửi thành công")
}
