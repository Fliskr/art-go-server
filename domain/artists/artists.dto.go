package artists

import (
	"errors"
	"log"
	"strconv"

	"github.com/gervi/art-go-server/client/sql_lite"
	"github.com/gervi/art-go-server/domain/artworks"
)

type Artist struct {
	ID       string              `json:"id"`
	Name     string              `json:"name"`
	Artworks []*artworks.Artwork `json:"artworks,omitempty"`
}

const (
	queryGet            = "SELECT * from artists where id=?;"
	queryInsertArtist   = "INSERT INTO artists(Name) values(?);"
	queryUpdateArtist   = "UPDATE artists set name=? where id=?;"
	queryDeleteArtist   = "DELETE from artists where id=?;"
	queryGetAll         = "SELECT * from artists;"
	queryByName         = "SELECT * from artists where name like ?;"
	queryGetArtworks    = "SELECT id,title from artworks where artist_id=?;"
	queryDeleteArtworks = "DELETE from artworks where artist_id=?;"
)

func (a *Artist) Get() {
	err := sql_lite.Client.QueryRow(queryGet, &a.ID).Scan(&a.ID, &a.Name)
	a.GetArtworks()
	if err != nil {
		if err.Error() != "sql: no rows in result set" {
			panic(err.Error())
		}
	}
}

func (a *Artist) Save() {
	result, err := sql_lite.Client.Exec(queryInsertArtist, &a.Name)
	if err != nil {
		panic(err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}
	log.Println("artist saved")
	a.ID = strconv.Itoa(int(id))
}

func (a *Artist) Update() error {
	result, err := sql_lite.Client.Exec(queryUpdateArtist, &a.Name, &a.ID)
	if err != nil {
		return err
	}
	rows, _ := result.RowsAffected()
	if rows == 0 {
		return errors.New("No artist with such ID")
	}
	log.Println("artist updated")
	return nil
}

func (a *Artist) Delete() error {
	_, err := sql_lite.Client.Exec(queryDeleteArtist, &a.ID)
	if err != nil {
		return err
	}
	log.Println("artist deleted")
	_, err = sql_lite.Client.Exec(queryDeleteArtworks, &a.ID)
	if err != nil {
		return err
	}
	log.Println("artist's artworks deleted")
	return nil
}

func (a *Artist) GetArtworks() {
	id, err := strconv.ParseInt(a.ID, 10, 64)
	rows, err := sql_lite.Client.Query(queryGetArtworks, id)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var artwork artworks.Artwork
		err := rows.Scan(&artwork.ID, &artwork.Title)
		if err != nil {
			panic(err)
		}
		a.Artworks = append(a.Artworks, &artwork)
	}
}

type Artists []Artist

func (a *Artists) GetAll() {
	rows, err := sql_lite.Client.Query(queryGetAll)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var artist Artist
		err := rows.Scan(&artist.ID, &artist.Name)
		if err != nil {
			panic(err)
		}
		artist.GetArtworks()
		*a = append(*a, artist)
	}
}

func (a *Artists) GetArtistByName(name string) {
	rows, err := sql_lite.Client.Query(queryByName, "%"+name+"%")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var artist Artist
		err := rows.Scan(&artist.ID, &artist.Name)
		if err != nil {
			panic(err)
		}
		artist.GetArtworks()
		*a = append(*a, artist)
	}
}
