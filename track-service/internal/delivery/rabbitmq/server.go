package rabbitmq

import (
	rabbit "artist-service/pkg/client/rabbitmq"
	"encoding/json"
	"fmt"
	"log"
	"track-service/internal/domain"

	amqp "github.com/rabbitmq/amqp091-go"
)

type AMPQServer struct {
	conn     *amqp.Connection
	amqpChan *amqp.Channel
	us       domain.TrackUseCase
}

func New(url string, us domain.TrackUseCase) (*AMPQServer, error) {
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
		us:       us}, nil
}

func (c *AMPQServer) StartConsumer() {
	q, err := c.amqpChan.QueueDeclare(
		"track",
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
			tracks := c.CreateTrack(track)

			for _, tag := range tracks.Tags {

				id, err := c.CreateTag(domain.CreateTagDTO{
					Name: tag.Name,
					URL:  &tag.URL,
				})

				if err != nil {
					return
				}

				if err := c.CreateTrackTag(id, *tracks.ID); err != nil {
					return
				}
			}
		}
	}
}

func (c *AMPQServer) CreateTrack(track domain.Track) *domain.Track {
	tracks, err := c.us.CreateTrack(track)
	if err != nil {
		return nil
	}

	return tracks
}

func (c *AMPQServer) CreateTag(tag domain.CreateTagDTO) (int, error) {
	id, err := c.us.CreateTag(tag)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (c *AMPQServer) CreateTrackTag(tagID int, trackID int) error {
	err := c.us.CreateTrackTag(tagID, trackID)
	if err != nil {
		return err
	}
	return nil
}
