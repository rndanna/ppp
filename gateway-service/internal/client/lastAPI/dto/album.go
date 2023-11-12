package dto

import (
	"gateway-service/internal/domain"
)

type Album struct {
	Title string `json:"title"`
	URL   string `json:"url"`
}

type AlbumSearch struct {
	Album struct {
		URL    string `json:"url"`
		Name   string `json:"name"`
		Tracks struct {
			Track []struct {
				Name   string `json:"name"`
				URL    string `json:"url"`
				Artist struct {
					Name string `json:"name"`
				} `json:"artist"`
			} `json:"track"`
		} `json:"tracks"`
		Tags struct {
			Tags []Tag `json:"tag"`
		} `json:"tags"`
	} `json:"album"`
}

func (t *AlbumSearch) ToDomain() ([]*domain.Track, error) {
	var tags []*domain.Tag

	for i := range t.Album.Tracks.Track {
		tags = append(tags, &domain.Tag{
			Name: t.Album.Tracks.Track[i].Name,
			URL:  t.Album.Tracks.Track[i].URL,
		})
	}

	var tracks []*domain.Track

	for i := range t.Album.Tracks.Track {
		track := domain.Track{
			Name: t.Album.Tracks.Track[i].Name,
			URL:  &t.Album.Tracks.Track[i].URL,
			Album: &domain.Album{
				Title: t.Album.Name,
				URL:   &t.Album.URL,
				Artist: &domain.Artist{
					Name: t.Album.Tracks.Track[i].Artist.Name,
				},
			},
			Tags: tags,
		}

		tracks = append(tracks, &track)
	}

	return tracks, nil
}
