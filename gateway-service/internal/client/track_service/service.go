package track_service

import (
	"encoding/json"
	"fmt"
	"gateway-service/internal/domain"
	"gateway-service/internal/usecase"
	"gateway-service/pkg/api"
	"io"
	"net/http"
	"time"
)

type client struct {
	base api.BaseClient
}

func New() usecase.TrackService {
	return &client{
		base: api.BaseClient{
			HTTPClient: &http.Client{
				Timeout: 10 * time.Second,
			},
		},
	}
}

func (c *client) GetTracksByTag(id int) ([]*domain.Track, error) {
	url := fmt.Sprintf("http://localhost:8083/track/tag/%d", id)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create new request due to error: %w", err)
	}

	response, err := c.base.SendRequest(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request due to error: %w", err)
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("SearchTrack ReadAll: %w", err)
	}

	var track []*domain.Track

	err = json.Unmarshal(body, &track)
	if err != nil {
		return nil, fmt.Errorf("Unmarshal err: %w", err)
	}

	return track, nil
}

func (c *client) GetTracksByArtist(id int) ([]*domain.Track, error) {
	url := fmt.Sprintf("http://localhost:8083/track/artist/%d", id)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create new request due to error: %w", err)
	}

	response, err := c.base.SendRequest(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request due to error: %w", err)
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("SearchTrack ReadAll: %w", err)
	}

	var track []*domain.Track

	err = json.Unmarshal(body, &track)
	if err != nil {
		return nil, fmt.Errorf("Unmarshal err: %w", err)
	}

	return track, nil
}

func (c *client) GetTracksChart(chart string) ([]*domain.Track, error) {
	url := fmt.Sprintf("http://localhost:8083/track/chart?chart=%s", chart)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create new request due to error: %w", err)
	}

	response, err := c.base.SendRequest(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request due to error: %w", err)
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("SearchTrack ReadAll: %w", err)
	}

	var track []*domain.Track

	err = json.Unmarshal(body, &track)
	if err != nil {
		return nil, fmt.Errorf("Unmarshal err: %w", err)
	}

	return track, nil
}
