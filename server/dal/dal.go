package dal

import (
	"database/sql"
	"fmt"
	"log"
	"recordings/model"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

func init() {
	// Capture connection properties.
	cfg := mysql.Config{
		// User:   os.Getenv("DBUSER"),
		// Passwd: os.Getenv("DBPASS"),
		User:                 "root",
		Passwd:               "root",
		Net:                  "tcp",
		Addr:                 "127.0.0.1:3306",
		DBName:               "recordings",
		AllowNativePasswords: true,
	}
	// Get a database handle.
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected to the database!")
}

func GetAll() ([]model.Album, error) {
	var albums []model.Album

	rows, err := db.Query("SELECT * FROM albums")
	if err != nil {
		defer rows.Close()
	}
	if err == nil {
		// Loop through rows, using Scan to assign column data to struct fields.
		for rows.Next() {
			var alb model.Album
			err = rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price)
			albums = append(albums, alb)
		}
	}
	return albums, err
}

// AlbumsByArtist queries for albums that have the specified artist name.
func AlbumsByArtist(name string) ([]model.Album, error) {
	// An albums slice to hold data from returned rows.
	var albums []model.Album

	rows, err := db.Query("SELECT * FROM albums WHERE artist = ?", name)
	if err != nil {
		return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
	}
	defer rows.Close()
	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var alb model.Album
		if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
			return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
		}
		albums = append(albums, alb)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
	}
	return albums, nil
}

// AlbumByID returns album by id
func AlbumByID(id int64) (model.Album, error) {
	var alb model.Album
	row := db.QueryRow("SELECT * FROM albums WHERE id = ? ", id)
	if err := row.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
		if err == sql.ErrNoRows {
			return alb, fmt.Errorf("albumsById %d: no such album", id)
		}
		return alb, fmt.Errorf("albumByID %q: %v", id, err)
	}
	return alb, nil
}

// AddAlbum adds a new album to the db
func AddAlbum(alb model.Album) (int64, error) {
	result, err := db.Exec("INSERT INTO albums (title, artist, price) VALUES (?, ?, ?) ", alb.Title, alb.Artist, alb.Price)
	if err != nil {
		return 0, fmt.Errorf("addAlbum: %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("addAlbum: %v", err)
	}
	return id, nil
}
