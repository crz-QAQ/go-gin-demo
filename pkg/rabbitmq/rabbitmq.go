package mq

import (
	"log"
	"os"

	"github.com/streadway/amqp"
)

var (
	conn *amqp.Connection
	ch   *amqp.Channel
)

// 初始化MQ
func Init() error {
	// 优先读线上环境变量
	mqUrl := os.Getenv("RABBITMQ_URL")
	if mqUrl == "" {
		// 本地默认地址
		mqUrl = "amqp://guest:guest@127.0.0.1:5672/"
	}
	var err error
	conn, err = amqp.Dial(mqUrl)
	if err != nil {
		return err
	}
	ch, err = conn.Channel()
	if err != nil {
		return err
	}
	log.Println("RabbitMQ 初始化连接成功")
	return nil
}

// 声明队列
func DeclareQueue(queueName string) (amqp.Queue, error) {
	return ch.QueueDeclare(
		queueName,
		true, // 持久化
		false,
		false,
		false,
		nil,
	)
}

// 发送消息
func SendMsg(queue string, body []byte) error {
	return ch.Publish(
		"",
		queue,
		false,
		false,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			Body:         body,
		},
	)
}

// 消费消息
func Consume(queue string, handle func([]byte)) {
	msgs, err := ch.Consume(queue, "", true, false, false, false, nil)
	if err != nil {
		log.Println("消费失败", err)
		return
	}
	go func() {
		for m := range msgs {
			handle(m.Body)
		}
	}()
}

// 关闭
func Close() {
	_ = ch.Close()
	_ = conn.Close()
}
