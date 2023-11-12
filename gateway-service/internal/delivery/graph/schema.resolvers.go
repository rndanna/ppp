package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.40

import (
	"context"
	"gateway-service/internal/domain"
)

// SearchTrack is the resolver for the searchTrack field.
func (r *queryResolver) SearchTrack(ctx context.Context, artist string, track string) (*domain.Track, error) {
	t, err := r.gatewayUS.SearchTrack(artist, track)
	if err != nil {
		return nil, err
	}

	return t, nil
}

// SearchAlbum is the resolver for the searchAlbum field.
func (r *queryResolver) SearchAlbum(ctx context.Context, artist string, album string) ([]*domain.Track, error) {
	t, err := r.gatewayUS.SearchAlbum(artist, album)
	if err != nil {
		return nil, err
	}

	return t, nil
}

// GetTracksByTag is the resolver for the getTracksByTag field.
func (r *queryResolver) GetTracksByTag(ctx context.Context, id int) ([]*domain.Track, error) {
	t, err := r.gatewayUS.GetTracksByTag(id)
	if err != nil {
		return nil, err
	}

	return t, nil
}

// GetTracksArtist is the resolver for the getTracksArtist field.
func (r *queryResolver) GetTracksArtist(ctx context.Context, id int) ([]*domain.Track, error) {
	t, err := r.gatewayUS.GetTracksByArtist(id)
	if err != nil {
		return nil, err
	}

	return t, nil
}

// GetChart is the resolver for the getChart field.
func (r *queryResolver) GetChart(ctx context.Context, chart string) ([]*domain.Track, error) {
	t, err := r.gatewayUS.GetTracksChart(chart)
	if err != nil {
		return nil, err
	}

	return t, nil
}

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
