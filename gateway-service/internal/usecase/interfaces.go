package usecase

import (
	"gateway-service/internal/domain"
)

type AMPQServer interface {
	Publish(queueName string, body []byte, types string) error
}

type LastFMService interface {
	SearchTrack(artist string, track string) (*domain.Track, error)
	SearchAlbum(artist string, album string) ([]*domain.Track, error)
}

type ArtistService interface {
	GetArtist(id int) (*domain.Artist, error)
}

type TrackService interface {
	GetTracksByTag(id int) ([]*domain.Track, error)
	GetTracksByArtist(id int) ([]*domain.Track, error)
	GetTracksChart(chart string) ([]*domain.Track, error)
}

type GatewayUSI interface {
	SearchTrack(artist string, track string) (*domain.Track, error)
	SearchAlbum(artist string, album string) ([]*domain.Track, error)
	GetTracksByTag(id int) ([]*domain.Track, error)
	GetTracksByArtist(id int) ([]*domain.Track, error)
	GetTracksChart(chart string) ([]*domain.Track, error)
}
