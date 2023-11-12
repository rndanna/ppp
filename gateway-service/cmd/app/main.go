package main

import (
	"fmt"
	"gateway-service/internal/client/artist_service"
	lastapi_service "gateway-service/internal/client/lastAPI"
	"gateway-service/internal/client/rabbitmq"
	"gateway-service/internal/client/track_service"
	"gateway-service/internal/delivery/graph"

	"gateway-service/pkg/utils/config"

	"gateway-service/internal/usecase"
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

func main() {
	cfg, err := config.Init()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}
	dsn := fmt.Sprintf(
		"amqp://%s:%s@%s:%s/",
		cfg.RabbitMQ.Login,
		cfg.RabbitMQ.Password,
		cfg.RabbitMQ.Host,
		cfg.RabbitMQ.Port,
	)
	fmt.Println(dsn)

	apiClient := lastapi_service.New("f95fe292baff414000911645bf2ba1c0")
	artistClient := artist_service.New()
	trackClient := track_service.New()
	publisher, err := rabbitmq.New(dsn)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	publisher.StartConsumer()

	gatewayUS := usecase.New(publisher, apiClient, artistClient, trackClient)
	res := graph.NewResolver(gatewayUS)
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: res}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", "8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
