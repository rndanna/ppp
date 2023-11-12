package usecase

import (
	"album-service/internal/domain"
	"fmt"
)

type AlbumUseCase struct {
	repo domain.AlbumRepository
}

func New(repo domain.AlbumRepository) domain.AlbumUseCase {
	return &AlbumUseCase{repo}
}

func (u *AlbumUseCase) GetAlbumByTitle(title string) (*domain.Album, error) {
	artist, err := u.repo.GetAlbumByTitle(title)
	if err != nil {
		return nil, fmt.Errorf("err UseCase - GetAlbumByTitle: %s", err)
	}

	return artist, nil
}

func (u *AlbumUseCase) CreateAlbum(track domain.Track) (*domain.Track, error) {
	alb, err := u.repo.GetAlbumByTitle(track.Album.Title)
	if err != nil || alb == nil {
		id, err := u.repo.CreateAlbum(domain.CreateAlbumDTO{
			Title:    track.Album.Title,
			URL:      *track.Album.URL,
			ArtistID: *track.Album.Artist.ID,
		})
		if err != nil {
			return nil, fmt.Errorf("err UseCase - CreateAlbum: %s", err)
		}

		track.Album.ID = &id
		return &track, nil
	}

	track.Album.ID = alb.ID

	return &track, nil
}
