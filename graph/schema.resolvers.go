package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"

	"github.com/gervi/art-go-server/domain/artists"
	"github.com/gervi/art-go-server/domain/artworks"
	"github.com/gervi/art-go-server/graph/generated"
	"github.com/gervi/art-go-server/graph/model"
)

func (r *mutationResolver) CreateArtist(ctx context.Context, input model.NewArtist) (*model.Artist, error) {
	artist := &artists.Artist{
		Name: input.Name,
	}
	artist.Save()
	artistModel := &model.Artist{
		ID:   artist.ID,
		Name: artist.Name,
	}
	return artistModel, nil
}

func (r *mutationResolver) CreateArtwork(ctx context.Context, input model.NewArtwork) (*model.Artwork, error) {
	artwork := &artworks.Artwork{Title: input.Title, ArtistId: input.Artist}
	err := artwork.Save()
	if err != nil {
		panic(err)
	}
	artist := &artists.Artist{ID: artwork.ArtistId}
	artist.Get()
	return &model.Artwork{ID: artwork.ID, Title: artwork.Title, Artist: &model.Artist{ID: artist.ID, Name: artist.Name}}, nil
}

func (r *mutationResolver) UpdateArtist(ctx context.Context, input model.NewArtist) (*model.Artist, error) {
	artist := &artists.Artist{ID: *input.ID, Name: input.Name}
	err := artist.Update()
	if err != nil {
		return nil, err
	}
	return &model.Artist{ID: artist.ID, Name: artist.Name}, nil
}

func (r *mutationResolver) UpdateArtwork(ctx context.Context, input model.NewArtwork) (*model.Artwork, error) {
	artwork := &artworks.Artwork{ID: *input.ID, Title: input.Title, ArtistId: input.Artist}
	err := artwork.Update()
	if err != nil {
		panic(err)
	}
	artist := &artists.Artist{ID: artwork.ArtistId}
	artist.Get()
	return &model.Artwork{ID: artwork.ID, Title: artwork.Title, Artist: &model.Artist{ID: artist.ID, Name: artist.Name}}, nil
}

func (r *mutationResolver) DeleteArtist(ctx context.Context, input string) (string, error) {
	artist := &artists.Artist{ID: input}
	err := artist.Delete()
	if err != nil {
		return "", err
	}
	return "success", nil
}

func (r *mutationResolver) DeleteArtwork(ctx context.Context, input string) (string, error) {
	artwork := &artworks.Artwork{ID: input}
	err := artwork.Delete()
	if err != nil {
		return "", err
	}
	return "success", nil
}

func (r *queryResolver) Artists(ctx context.Context, name *string) ([]*model.Artist, error) {
	var artists artists.Artists
	if *name == "" {
		artists.GetAll()
	} else {
		artists.GetArtistByName(*name)
	}
	var modelArtists = make([]*model.Artist, len(artists))
	for i, artist := range artists {
		var arts = make([]*model.Artwork, len(artist.Artworks))
		for i, art := range artist.Artworks {
			arts[i] = &model.Artwork{ID: art.ID, Title: art.Title}
		}
		modelArtists[i] = &model.Artist{ID: artist.ID, Name: artist.Name, Artworks: arts}
	}
	return modelArtists, nil
}

func (r *queryResolver) Artworks(ctx context.Context, artist *string) ([]*model.Artwork, error) {
	if *artist == "0" || *artist == "" {
		return nil, errors.New("artist should be not null")
	}
	artist_m := &artists.Artist{ID: *artist}
	artist_m.GetArtworks()
	artworks := make([]*model.Artwork, len(artist_m.Artworks))
	for i, artwork := range artist_m.Artworks {
		artworks[i] = &model.Artwork{ID: artwork.ID, Title: artwork.Title}
	}
	return artworks, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
