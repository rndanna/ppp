package app

import (
	albumHttpDelivery "album-service/internal/delivery/http"
	"album-service/internal/delivery/rabbitmq"
	"album-service/internal/repository/postgresql"
	"album-service/internal/usecase"
	"album-service/pkg/client/postgres"
	"album-service/pkg/utils/config"
	"fmt"
	"log"

	"github.com/labstack/echo/v4"
)

func Run(cfg *config.Config) {
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.Database.Postgresql.Username,
		cfg.Database.Postgresql.Password,
		cfg.Database.Postgresql.Host,
		cfg.Database.Postgresql.Port,
		cfg.Database.Postgresql.Name,
	)

	pg, err := postgres.New(dsn)
	if err != nil {
		log.Fatalf("postgres New: %v", err)
	}
	e := echo.New()
	dsn = fmt.Sprintf(
		"amqp://%s:%s@%s:%s/",
		cfg.RabbitMQ.Login,
		cfg.RabbitMQ.Password,
		cfg.RabbitMQ.Host,
		cfg.RabbitMQ.Port,
	)

	repo := postgresql.New(pg)
	us := usecase.New(repo)

	albumHttpDelivery.New(e, us)
	r, err := rabbitmq.New(dsn, us)
	if err != nil {
		return
	}

	r.StartConsumer()

	log.Fatal(e.Start(fmt.Sprintf(":%s", cfg.Server.Port)))
}
