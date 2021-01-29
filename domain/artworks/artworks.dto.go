package artworks

import (
	"errors"
	"log"
	"strconv"

	"github.com/gervi/art-go-server/client/sql_lite"
)

const (
	queryDeleteByArtistId = `DELETE from artworks where artist_id=?;`
	queryDelete           = `DELETE from artworks where id=?;`
	queryInsert           = "INSERT INTO artworks(title, artist_id) values(?,?);"
	queryUpdate           = "UPDATE artworks set title=?, artist_id=? where id=?;"
)

type Artwork struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	ArtistId string `json:"artist_id"`
}

func (a *Artwork) DeleteByArtistId(id string) error {
	_, err := sql_lite.Client.Exec(queryDeleteByArtistId, id)
	if err != nil {
		return err
	}
	log.Println("arworks deleted")
	return nil
}

func (a *Artwork) Delete() error {
	_, err := sql_lite.Client.Exec(queryDelete, a.ID)
	if err != nil {
		return err
	}
	log.Println("artwork deleted")
	return nil
}

func (a *Artwork) Save() error {
	result, err := sql_lite.Client.Exec(queryInsert, &a.Title, &a.ArtistId)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return errors.New("Artwork wasn't created")
	}
	a.ID = strconv.Itoa(int(id))
	log.Println("artwork saved")
	return nil
}

func (a *Artwork) Update() error {
	_, err := sql_lite.Client.Exec(queryUpdate, &a.Title, &a.ArtistId, &a.ID)
	if err != nil {
		return err
	}
	log.Println("artwork updated")
	return nil
}

type Artworks []Artwork
