package lastapi_service

import (
	"encoding/json"
	"fmt"
	"gateway-service/internal/client/lastAPI/dto"
	"gateway-service/internal/domain"
	"gateway-service/internal/usecase"
	"gateway-service/pkg/api"
	"io"
	"net/http"
	"net/url"
	"time"
)

type client struct {
	base      api.BaseClient
	keySecret string
}

func New(keySecret string) usecase.LastFMService {
	return &client{
		base: api.BaseClient{
			HTTPClient: &http.Client{
				Timeout: 10 * time.Second,
			},
		},
		keySecret: keySecret,
	}
}

const (
	BaseURL = "https://ws.audioscrobbler.com/2.0/"
)

func (c *client) SearchTrack(artist string, track string) (*domain.Track, error) {
	// url := fmt.Sprintf("https://ws.audioscrobbler.com/2.0/?method=track.getInfo&api_key=f95fe292baff414000911645bf2ba1c0&artist=%s&track=%s&format=json", artist, track)

	params := url.Values{}
	params.Add("method", "track.getInfo")
	params.Add("api_key", c.keySecret)
	params.Add("artist", artist)
	params.Add("track", track)
	params.Add("format", "json")

	// var dto dto.TrackSearch

	resp, err := c.makeRequest(http.MethodGet, params)
	if err != nil {
		return nil, fmt.Errorf("make request: %w", err)
	}

	t, err := resp.TrackSearch.ToDomain()
	if err != nil {
		return nil, fmt.Errorf("toDomain err: %w", err)
	}

	fmt.Println(t)

	return t, nil
}

func (c *client) SearchAlbum(artist string, album string) ([]*domain.Track, error) {
	url := fmt.Sprintf("https://ws.audioscrobbler.com//2.0/?method=album.getinfo&api_key=f95fe292baff414000911645bf2ba1c0&artist=%s&album=%s&format=json", artist, album)

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
		return nil, fmt.Errorf("SearchAlbum ReadAll: %w", err)
	}

	var dto dto.AlbumSearch

	err = json.Unmarshal(body, &dto)
	if err != nil {
		return nil, err
	}

	t, err := dto.ToDomain()
	if err != nil {
		return nil, fmt.Errorf("toDomain err: %w", err)
	}

	return t, nil
}

func (c *client) makeRequest(method string, params url.Values) (*LastFM, error) {
	req, err := http.NewRequest(method, BaseURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create new request due to error: %w", err)
	}
	req.URL.RawQuery = params.Encode()

	response, err := c.base.SendRequest(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request due to error: %w", err)
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("ReadAll: %w", err)
	}
	defer response.Body.Close()

	var lastfm LastFM
	err = json.Unmarshal(body, &lastfm)
	if err != nil {
		return nil, fmt.Errorf("Unmarshal err: %w", err)
	}

	return &lastfm, nil
}
