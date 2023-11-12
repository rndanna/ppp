package domain

type TrackRepository interface {
	GetTrack(id int) (*Track, error)
	CreateTrack(dto CreateTrackDTO) (int, error)
	GetTrackByTag(id int) ([]*Track, error)
	GetTrackByArtist(id int) ([]*Track, error)
	GetTopTrackByListeners() ([]*Track, error)
	GetTopTrackByPlayCount() ([]*Track, error)
	CreateTag(dto CreateTagDTO) (int, error)
	CreateTrackTag(tagID, trackID int) error
}

type TrackUseCase interface {
	GetTrack(id int) (*Track, error)
	CreateTrack(track Track) (*Track, error)
	GetTrackByTag(id int) ([]*Track, error)
	GetTrackByArtist(id int) ([]*Track, error)
	GetTopTrackByListeners() ([]*Track, error)
	GetTopTrackByPlayCount() ([]*Track, error)
	CreateTag(tag CreateTagDTO) (int, error)
	CreateTrackTag(tagID, trackID int) error
}

//TODO context
