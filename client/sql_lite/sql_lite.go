package sql_lite

import (
	"database/sql"
	"flag"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

var (
	Client *sql.DB
)

const (
	createArtistsTable = `CREATE TABLE IF NOT EXISTS artists (
		id integer not null primary key,
		name VARCHAR (127) NOT NULL UNIQUE
	);`
	createArtworksTable = `CREATE TABLE IF NOT EXISTS artworks (
		id integer not null primary key,
		title VARCHAR (127) NOT NULL UNIQUE,
		artist_id INT,
		FOREIGN KEY (artist_id) REFERENCES artists(ID)
	);`
	createArtist  = `insert into artists(name) values(?);`
	createArtwork = `insert into artworks(title, artist_id) values(?, ?);`
)

func init() {
	os.Remove("./art.db")
	var err error
	Client, err = sql.Open("sqlite3", "./art.db")
	if err != nil {
		panic(err)
	}
	if err = Client.Ping(); err != nil {
		panic(err)
	}
	log.Println("Connection established")

	createTable()
	var fillSample = flag.Bool("fill", false, "help message for flag n")
	flag.Parse()
	if *fillSample {
		fillTableWithSampleData()
	}
}

func createTable() {
	_, err := Client.Exec(createArtistsTable)
	if err != nil {
		log.Println(err.Error())
		panic(err)
	}
	_, err = Client.Exec(createArtworksTable)
	if err != nil {
		panic(err)
	}
	log.Println("Databases created")
}

func fillTableWithSampleData() {
	artists := []string{"Leonardo da Vinci", "Michelangelo", "Rembrandt van Rijn", "Vincent van Gogh"}
	artworks := []string{"Mona Lisa", "The Creation of Adam", "The Anatomy Lesson of Dr Nicolaes Tulp", "Starry Night"}
	artworks2 := []string{"Lisa Moana", "Adam's Creation"}
	for _, artist := range artists {
		Client.Exec(createArtist, artist)
	}
	for id, artwork := range artworks {
		Client.Exec(createArtwork, artwork, id+1)
	}
	for id, artwork := range artworks2 {
		Client.Exec(createArtwork, artwork, id+1)
	}
	log.Println("Sample data filled")
}
