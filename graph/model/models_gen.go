// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Artist struct {
	ID       string     `json:"id"`
	Name     string     `json:"name"`
	Artworks []*Artwork `json:"artworks"`
}

type Artwork struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist *Artist `json:"artist"`
}

type NewArtist struct {
	Name string  `json:"name"`
	ID   *string `json:"id"`
}

type NewArtwork struct {
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	ID     *string `json:"id"`
}
