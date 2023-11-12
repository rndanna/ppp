package dto

import "gateway-service/internal/domain"

type Artist struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type ArtistCreateDTO struct {
	Name    string `json:"name"`
	URL     string `json:"url"`
	Content string `json:"content"`
}

func (a *Artist) toDomain() *domain.Artist {
	art := domain.Artist{
		Name: a.Name,
		URL:  &a.URL,
	}
	return &art
}
