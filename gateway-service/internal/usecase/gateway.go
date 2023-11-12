package usecase

import (
	"encoding/json"
	"fmt"
	"gateway-service/internal/domain"
)

type GatewayUS struct {
	ampq          AMPQServer
	lastAPI       LastFMService
	ArtistService ArtistService
	TrackService  TrackService
}

func New(ampq AMPQServer, lastAPI LastFMService, artServ ArtistService, TrackService TrackService) GatewayUSI {
	return &GatewayUS{
		ampq:          ampq,
		lastAPI:       lastAPI,
		ArtistService: artServ,
		TrackService:  TrackService,
	}
}

func (us *GatewayUS) SearchTrack(artist string, track string) (*domain.Track, error) {
	t, err := us.lastAPI.SearchTrack(artist, track)
	if err != nil {
		return nil, fmt.Errorf("SearchTrack err: %w", err)
	}

	b, err := json.Marshal(t)
	if err != nil {
		return nil, fmt.Errorf("Marshal err: %w", err)
	}

	us.ampq.Publish("artist", b, "create")

	return t, nil
}

func (us *GatewayUS) SearchAlbum(artist string, album string) ([]*domain.Track, error) {
	t, err := us.lastAPI.SearchAlbum(artist, album)
	if err != nil {
		return nil, fmt.Errorf("SearchAlbum err: %w", err)
	}

	b, err := json.Marshal(t)
	if err != nil {
		return nil, fmt.Errorf("Marshal err: %w", err)
	}

	us.ampq.Publish("artist", b, "create")

	return t, nil
}

func (us *GatewayUS) GetTracksByTag(id int) ([]*domain.Track, error) {
	t, err := us.TrackService.GetTracksByTag(id)
	if err != nil {
		return nil, fmt.Errorf("GetTracksByTag err: %w", err)
	}

	return t, nil
}

func (us *GatewayUS) GetTracksByArtist(id int) ([]*domain.Track, error) {
	t, err := us.TrackService.GetTracksByArtist(id)
	if err != nil {
		return nil, fmt.Errorf("GetTracksByArtist err: %w", err)
	}

	return t, nil
}
func (us *GatewayUS) GetTracksChart(chart string) ([]*domain.Track, error) {
	t, err := us.TrackService.GetTracksChart(chart)
	if err != nil {
		return nil, fmt.Errorf("GetTracksChart err: %w", err)
	}

	return t, nil
}
