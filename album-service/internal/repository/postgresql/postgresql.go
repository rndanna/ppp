package postgresql

import (
	"album-service/internal/domain"
	"album-service/pkg/client/postgres"
	"fmt"
)

type AlbumRepository struct {
	*postgres.Postgres
}

func New(pg *postgres.Postgres) domain.AlbumRepository {
	return &AlbumRepository{pg}
}

func (r *AlbumRepository) GetAlbumByTitle(title string) (*domain.Album, error) {
	var album domain.Album

	if err := r.Pool.QueryRow(`
		SELECT id, title, url
		FROM albums
		WHERE title = $1
	`, title).Scan(&album.ID, &album.Title, &album.URL); err != nil {

		return nil, fmt.Errorf("err GetAlbumByTitle QueryRow: %w", err)
	}

	return &album, nil
}

func (r *AlbumRepository) CreateAlbum(dto domain.CreateAlbumDTO) (int, error) {
	var id int
	if err := r.Pool.QueryRow(`
		INSERT INTO  albums(title, url, artist_id)
		VALUES($1, $2, $3)
		RETURNING id
    `, dto.Title, dto.URL, dto.ArtistID).Scan(&id); err != nil {
		return 0, fmt.Errorf("err CreateAlbum Exec: %w", err)
	}

	return id, nil
}
