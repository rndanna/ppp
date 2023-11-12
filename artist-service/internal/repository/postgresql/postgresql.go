package postgresql

import (
	"artist-service/internal/domain"
	"artist-service/pkg/client/postgres"
	"fmt"
)

type ArtistRepository struct {
	*postgres.Postgres
}

func New(pg *postgres.Postgres) domain.ArtistRepository {
	return &ArtistRepository{pg}
}

func (r *ArtistRepository) GetArtist(id int) (*domain.Artist, error) {
	var artist domain.Artist

	if err := r.Pool.QueryRow(`
		SELECT id, name, url
		FROM artists
		WHERE id = $1
	`, id).Scan(&artist.ID, &artist.Name, &artist.URL); err != nil {

		return nil, fmt.Errorf("err GetArtist QueryRow: %w", err)
	}

	return &artist, nil
}

func (r *ArtistRepository) GetArtistByName(name string) (*domain.Artist, error) {
	var artist domain.Artist

	if err := r.Pool.QueryRow(`
		SELECT id, name, url
		FROM artists
		WHERE name = $1
	`, name).Scan(&artist.ID, &artist.Name, &artist.URL); err != nil {

		return nil, fmt.Errorf("err GetArtistByName QueryRow: %w", err)
	}

	return &artist, nil
}

func (r *ArtistRepository) CreateArtist(dto domain.CreateArtistDTO) (int, error) {
	var id int
	if err := r.Pool.QueryRow(`
		INSERT INTO  artists(name, url)
		VALUES($1, $2)
		RETURNING id
    `, dto.Name, dto.URL).Scan(&id); err != nil {
		return 0, fmt.Errorf("err CreateArtist Exec: %w", err)
	}

	return id, nil
}
