package usecase

import (
	"fmt"
	"track-service/internal/domain"
)

type TrackUseCase struct {
	repo domain.TrackRepository
}

func New(repo domain.TrackRepository) domain.TrackUseCase {
	return &TrackUseCase{repo}
}

func (u *TrackUseCase) GetTrack(id int) (*domain.Track, error) {
	artist, err := u.repo.GetTrack(id)
	if err != nil {
		return nil, fmt.Errorf("err UseCase - GetTrack: %s", err)
	}

	return artist, nil
}

func (u *TrackUseCase) CreateTrack(track domain.Track) (*domain.Track, error) {
	id, err := u.repo.CreateTrack(domain.CreateTrackDTO{
		Name:      track.Name,
		URL:       track.URL,
		Playcount: track.Playcount,
		Listeners: track.Listeners,
		AlbumID:   *track.Album.ID,
	})
	if err != nil {
		return nil, fmt.Errorf("err UseCase - CreateTrack: %s", err)
	}

	track.ID = &id

	return &track, nil
}

func (u *TrackUseCase) CreateTag(tag domain.CreateTagDTO) (int, error) {
	id, err := u.repo.CreateTag(tag)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	return id, nil
}

func (u *TrackUseCase) CreateTrackTag(tagID, trackID int) error {
	err := u.repo.CreateTrackTag(tagID, trackID)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (u *TrackUseCase) GetTrackByTag(id int) ([]*domain.Track, error) {
	tracks, err := u.repo.GetTrackByTag(id)
	if err != nil {
		return tracks, fmt.Errorf("err UseCase - GetTrackByTag: %s", err)
	}

	return tracks, nil
}

func (u *TrackUseCase) GetTrackByArtist(id int) ([]*domain.Track, error) {
	tracks, err := u.repo.GetTrackByArtist(id)
	if err != nil {
		return tracks, fmt.Errorf("err UseCase - GetTrackByArtist: %s", err)
	}

	return tracks, nil
}

func (u *TrackUseCase) GetTopTrackByListeners() ([]*domain.Track, error) {
	tracks, err := u.repo.GetTopTrackByListeners()
	if err != nil {
		return tracks, fmt.Errorf("err UseCase - GetTopTrackByListeners: %s", err)
	}

	return tracks, nil
}

func (u *TrackUseCase) GetTopTrackByPlayCount() ([]*domain.Track, error) {
	tracks, err := u.repo.GetTopTrackByPlayCount()
	if err != nil {
		return tracks, fmt.Errorf("err UseCase - GetTopTrackByDuration: %s", err)
	}

	return tracks, nil
}
