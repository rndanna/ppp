package rabbitmq

import (
	"context"
	"fmt"
	rabbit "gateway-service/pkg/client/rabbitmq"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Server struct {
	conn     *amqp.Connection
	amqpChan *amqp.Channel
}

func New(url string) (*Server, error) {
	mqConn, err := rabbit.New(url)
	if err != nil {
		return nil, err
	}

	amqpChan, err := mqConn.Channel()
	if err != nil {
		return nil, fmt.Errorf("error Channel %w", err)
	}

	return &Server{
		conn:     mqConn,
		amqpChan: amqpChan}, nil
}

func (p *Server) Publish(queueName string, body []byte, types string) error {
	q, err := p.amqpChan.QueueDeclare(
		queueName,
		false,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		log.Fatalf("failed to declare a queue. Error: %s", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = p.amqpChan.PublishWithContext(ctx,
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
			Type:        types,
		})
	if err != nil {
		log.Fatalf("failed to publish a message. Error: %s", err)
		return nil
	}

	return nil
}

func (c *Server) StartConsumer() {
	q, err := c.amqpChan.QueueDeclare(
		"gw",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("failed to declare a queue. Error: %s", err)
	}

	messages, err := c.amqpChan.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("failed to register a consumer. Error: %s", err)
	}

	go c.worker(messages)
}

func (c *Server) worker(messages <-chan amqp.Delivery) {
	for delivery := range messages {
		switch delivery.Type {
		case "artist":
			c.Publish("album", delivery.Body, "create")
		case "album":
			c.Publish("track", delivery.Body, "create")
		}
	}
}
