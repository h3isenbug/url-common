package consumer

import (
	"sync"

	mux "github.com/h3isenbug/url-common/pkg/event-mux"
	"github.com/streadway/amqp"
)

type MessageQueueConsumer interface {
	ConsumeMessages() error
	GracefulShutdown() error
}

type RabbitMQConsumerV1 struct {
	messageMux mux.MessageMux
	mqChannel  *amqp.Channel
	ack        func(tag uint64)

	queue string

	wg *sync.WaitGroup
}

func NewRabbitMQQueueConsumerV1(
	mqChannel *amqp.Channel,
	ack func(tag uint64),
	messageMux mux.MessageMux,
	queueName string,
) (MessageQueueConsumer, error) {
	queue, err := mqChannel.QueueDeclare(queueName, true, false, true, false, nil)
	if err != nil {
		return nil, err
	}

	return &RabbitMQConsumerV1{
		messageMux: messageMux,
		mqChannel:  mqChannel,
		ack:        ack,
		queue:      queue.Name,
	}, nil
}

func (consumer RabbitMQConsumerV1) ConsumeMessages() error {
	var messages, err = consumer.mqChannel.Consume(
		consumer.queue,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	for message := range messages {
		consumer.messageMux.Handle(message.Body, func() { consumer.ack(message.DeliveryTag) })
	}

	return nil
}

func (consumer RabbitMQConsumerV1) GracefulShutdown() error {
	return consumer.mqChannel.Close()
}
