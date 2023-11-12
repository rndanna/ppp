package artist_service

import (
	"gateway-service/internal/domain"
	"gateway-service/internal/usecase"
	"gateway-service/pkg/api"
	"net/http"
	"time"
)

type client struct {
	base api.BaseClient
}

func New() usecase.ArtistService {
	return &client{
		base: api.BaseClient{
			HTTPClient: &http.Client{
				Timeout: 10 * time.Second,
			},
		},
	}
}

func (c *client) GetArtist(id int) (*domain.Artist, error) {
	return nil, nil
}
