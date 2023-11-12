package lastapi_service

import (
	"fmt"
	"gateway-service/internal/domain"
	"strconv"
)

type LastFM struct {
	TrackSearch Track       `json:"track"`
	AlbumSearch AlbumSearch `json:"album"`
}

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

type Tag struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Artist struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type ArtistCreateDTO struct {
	Name    string `json:"name"`
	URL     string `json:"url"`
	Content string `json:"content"`
}

// type TrackSearch struct {
// 	Track Track `json:"track"`
// }

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

func (a *Artist) toDomain() *domain.Artist {
	art := domain.Artist{
		Name: a.Name,
		URL:  &a.URL,
	}
	return &art
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
