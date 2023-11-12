package postgresql

import (
	"fmt"
	"track-service/internal/domain"
	"track-service/pkg/client/postgres"
)

type TrackRepository struct {
	*postgres.Postgres
}

func New(pg *postgres.Postgres) domain.TrackRepository {
	return &TrackRepository{pg}
}

func (r *TrackRepository) GetTrack(id int) (*domain.Track, error) {
	var track domain.Track

	if err := r.Pool.QueryRow(`
		SELECT id, name, url, playcount, listeners
		FROM tracks
		WHERE id = $1
	`, id).Scan(&track.ID, &track.URL, &track.Playcount, &track.Listeners); err != nil {

		return nil, fmt.Errorf("err GetAlbum QueryRow: %w", err)
	}

	return &track, nil
}

func (r *TrackRepository) CreateTrack(dto domain.CreateTrackDTO) (int, error) {
	var id int
	if err := r.Pool.QueryRow(`
		INSERT INTO  tracks(name, url, playcount, listeners, album_id)
		VALUES($1, $2, $3, $4, $5)
		RETURNING id
    `, dto.Name, dto.URL, dto.Playcount, dto.Listeners, dto.AlbumID).Scan(&id); err != nil {
		return 0, fmt.Errorf("err CreateTrack Exec: %w", err)
	}

	return id, nil
}

func (r *TrackRepository) CreateTag(dto domain.CreateTagDTO) (int, error) {
	var id int
	if err := r.Pool.QueryRow(`
		INSERT INTO  tags(name, url)
		VALUES($1, $2)
		RETURNING id
    `, dto.Name, dto.URL).Scan(&id); err != nil {
		return 0, fmt.Errorf("err CreateTag Exec: %w", err)
	}

	return id, nil
}

func (r *TrackRepository) CreateTrackTag(tagID, trackID int) error {
	var id int
	if err := r.Pool.QueryRow(`
		INSERT INTO  tracks_tag(tag_id, track_id)
		VALUES($1, $2)
		RETURNING id
    `, tagID, trackID).Scan(&id); err != nil {
		return fmt.Errorf("err CreateTag Exec: %w", err)
	}

	return nil
}

func (r *TrackRepository) GetTrackByTag(id int) ([]*domain.Track, error) {
	var tracks []*domain.Track

	rows, err := r.Pool.Query(`
		SELECT t.id, t.name, t.url, t.listeners, t.playcount, art.id, art.name
		FROM tracks_tag as tt
		JOIN tags  ON tt.tag_id = tags.id
		JOIN tracks as t ON tt.track_id = t.id
		JOIN albums as alb ON t.album_id = alb.id
		JOIN artists as art ON alb.artist_id = art.id
		WHERE tag_id = $1
	`, id)
	if err != nil {
		return tracks, fmt.Errorf("err GetTrackByTag Query: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var track domain.Track
		var art domain.Artist
		if err = rows.Scan(&track.ID, &track.Name, &track.URL, &track.Listeners, &track.Playcount, &art.ID, &art.Name); err != nil {
			return tracks, fmt.Errorf("err GetTrackByTag Scan: %w", err)
		}
		track.Artist = &art
		tracks = append(tracks, &track)
	}

	return tracks, nil
}

func (r *TrackRepository) GetTrackByArtist(id int) ([]*domain.Track, error) {
	var tracks []*domain.Track

	rows, err := r.Pool.Query(`
		SELECT t.id, t.name, t.url, t.listeners, t.playcount, art.id, art.name
		FROM tracks as t
		JOIN albums as alb ON t.album_id = alb.id
		JOIN artists as art ON alb.artist_id = art.id
		WHERE art.id = $1
	`, id)
	if err != nil {
		return tracks, fmt.Errorf("err GetTrackByArtist Query: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var track domain.Track
		var art domain.Artist
		if err = rows.Scan(&track.ID, &track.Name, &track.URL, &track.Listeners, &track.Playcount, &art.ID, &art.Name); err != nil {
			return tracks, fmt.Errorf("err GetTrackByArtist Scan: %w", err)
		}
		track.Artist = &art
		tracks = append(tracks, &track)
	}

	return tracks, nil
}

func (r *TrackRepository) GetTopTrackByListeners() ([]*domain.Track, error) {
	var tracks []*domain.Track

	rows, err := r.Pool.Query(`
		SELECT t.id, t.name, t.url, t.listeners, t.playcount, art.id, art.name
		FROM tracks as  t
		JOIN albums as alb ON t.album_id = alb.id
		JOIN artists as art ON alb.artist_id = art.id
		ORDER BY listeners DESC;
	`)
	if err != nil {
		return tracks, fmt.Errorf("err GetTopTrackByListeners Query: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var track domain.Track
		var art domain.Artist
		if err = rows.Scan(&track.ID, &track.Name, &track.URL, &track.Listeners, &track.Playcount, &art.ID, &art.Name); err != nil {
			return tracks, fmt.Errorf("err GetTopTrackByListeners Scan: %w", err)
		}
		track.Artist = &art
		tracks = append(tracks, &track)
	}

	return tracks, nil
}

func (r *TrackRepository) GetTopTrackByPlayCount() ([]*domain.Track, error) {
	var tracks []*domain.Track

	rows, err := r.Pool.Query(`
		SELECT t.id, t.name, t.url, t.listeners, t.playcount, art.id, art.name
		FROM tracks as  t
		JOIN albums as alb ON t.album_id = alb.id
		JOIN artists as art ON alb.artist_id = art.id
		ORDER BY playcount DESC;
	`)
	if err != nil {
		return tracks, fmt.Errorf("err GetTopTrackByPlayCount Query: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var track domain.Track
		var art domain.Artist
		if err = rows.Scan(&track.ID, &track.Name, &track.URL, &track.Listeners, &track.Playcount, &art.ID, &art.Name); err != nil {
			return tracks, fmt.Errorf("err GetTopTrackByPlayCount Scan: %w", err)
		}
		track.Artist = &art
		tracks = append(tracks, &track)
	}

	return tracks, nil
}
