package dto

import (
	"fmt"
	"gateway-service/internal/domain"
	"strconv"
)

type TrackSearch struct {
	Track Track `json:"track"`
}

type Track struct {
	Name      string `json:"name"`
	URL       string `json:"url"`
	Playcount string `json:"playcount"`
	Listeners string `json:"listeners"`
	Album     Album  `json:"album"`
	Artist    Artist `json:"artist"`
	TopTags   struct {
		Tags []Tag `json:"tag"`
	} `json:"toptags"`
}

func (t *Track) ToDomain() (*domain.Track, error) {
	playcount, err := strconv.Atoi(t.Playcount)
	if err != nil {
		return nil, fmt.Errorf("err strconv playcount: %w", err)
	}

	listeners, err := strconv.Atoi(t.Listeners)
	if err != nil {
		return nil, fmt.Errorf("err strconv listeners: %w", err)
	}

	var tags []*domain.Tag

	for i := range t.TopTags.Tags {
		tags = append(tags, &domain.Tag{
			Name: t.TopTags.Tags[i].Name,
			URL:  t.TopTags.Tags[i].URL,
		})
	}

	return &domain.Track{
		Name:      t.Name,
		URL:       &t.URL,
		Playcount: &playcount,
		Listeners: &listeners,
		Album: &domain.Album{
			Title: t.Album.Title,
			URL:   &t.Album.URL,
			Artist: &domain.Artist{
				Name: t.Artist.Name,
				URL:  &t.Artist.URL,
			},
		},
		Tags: tags,
	}, nil
}
