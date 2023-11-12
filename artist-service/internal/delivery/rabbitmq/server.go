package rabbitmq

import (
	"artist-service/internal/domain"
	rabbit "artist-service/pkg/client/rabbitmq"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

type AMPQServer struct {
	conn     *amqp.Connection
	amqpChan *amqp.Channel
	artistUS domain.ArtistUseCase
}

func New(url string, artistUS domain.ArtistUseCase) (*AMPQServer, error) {
	mqConn, err := rabbit.New(url)
	if err != nil {
		return nil, err
	}

	amqpChan, err := mqConn.Channel()
	if err != nil {
		return nil, fmt.Errorf("error Channel %w", err)
	}

	return &AMPQServer{
		conn:     mqConn,
		amqpChan: amqpChan,
		artistUS: artistUS}, nil
}

func (p *AMPQServer) Publish(queueName string, body []byte, types string) error {
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

func (c *AMPQServer) StartConsumer() {
	q, err := c.amqpChan.QueueDeclare(
		"artist",
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

func (c *AMPQServer) worker(messages <-chan amqp.Delivery) {
	for delivery := range messages {
		switch delivery.Type {
		case "create":
			var track domain.Track
			err := json.Unmarshal(delivery.Body, &track)
			if err != nil {
				return
			}
			c.CreateArtist(track)
		}
	}
}

func (c *AMPQServer) CreateArtist(track domain.Track) {
	t, err := c.artistUS.CreateArtist(track)
	if err != nil {
		return
	}
	b, err := json.Marshal(t)

	c.Publish("gw", b, "artist")
}
