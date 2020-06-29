package consumer

import (
	mux "github.com/h3isenbug/url-shortener/pkg/event-mux"
	"github.com/streadway/amqp"
	"sync"
)

type MessageQueueConsumer interface {
	ConsumeMessages() error
	GracefulShutdown() error
}

type RabbitMQConsumerV1 struct {
	messageMux mux.MessageMux
	shutdown   bool
	mqChannel  *amqp.Channel
	ack        func(tag uint64)

	queue string

	wg *sync.WaitGroup
}

func NewRabbitMQConsumerV1(
	mqChannel *amqp.Channel,
	ack func(tag uint64),
	messageMux mux.MessageMux,
	queue string,
) MessageQueueConsumer {
	return &RabbitMQConsumerV1{
		messageMux: messageMux,
		shutdown:   false,
		mqChannel:  mqChannel,
		ack:        ack,
		queue:      queue,
	}
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

	for !consumer.shutdown {
		var message = <-messages
		consumer.messageMux.Handle(message.Body, func() { consumer.ack(message.DeliveryTag) })
	}

	return nil
}

func (consumer RabbitMQConsumerV1) GracefulShutdown() error {
	consumer.shutdown = true
	return consumer.mqChannel.Close()
}
