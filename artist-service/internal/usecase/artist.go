package usecase

import (
	"artist-service/internal/domain"
	"fmt"
)

type ArtistUseCase struct {
	repo domain.ArtistRepository
}

func New(repo domain.ArtistRepository) domain.ArtistUseCase {
	return &ArtistUseCase{repo}
}

func (u *ArtistUseCase) GetArtist(id int) (*domain.Artist, error) {
	artist, err := u.repo.GetArtist(id)
	if err != nil {
		return nil, fmt.Errorf("err UseCase - GetArtist: %s", err)
	}

	return artist, nil
}

func (u *ArtistUseCase) GetArtistByName(name string) (*domain.Artist, error) {
	artist, err := u.repo.GetArtistByName(name)
	if err != nil {
		return nil, fmt.Errorf("err UseCase - GetArtistByName: %s", err)
	}

	return artist, nil
}

func (u *ArtistUseCase) CreateArtist(t domain.Track) (*domain.Track, error) {
	art, err := u.repo.GetArtistByName(t.Album.Artist.Name)
	if err != nil || art == nil {
		id, err := u.repo.CreateArtist(domain.CreateArtistDTO{
			Name: t.Album.Artist.Name,
			URL:  *t.Album.Artist.URL,
		})
		if err != nil {
			return nil, fmt.Errorf("err UseCase - CreateArtist: %s", err)
		}
		t.Album.Artist.ID = &id
		return &t, nil
	}

	t.Album.Artist.ID = art.ID
	return &t, nil
}
