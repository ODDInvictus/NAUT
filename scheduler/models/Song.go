package models

import (
	"strings"
)

type Song struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Artists string `json:"artist"`
	Album string `json:"album"`
	Cover string `json:"cover"`
	Duration int `json:"duration"`	
}

func GetSongByID(id string) (song, error) {
	song := []Song{}
	err := DB.Select(&song, "SELECT * FROM Song WHERE id = ?", id)

	if len(song) == 0 {
		return Song{}, err
	}

	return song[0], err
}

func GetSongs(ids string) ([]Song, error) {
	songs := []Song{}
	
	values, err := strings.Join(ids, ", ")

	if err != nil {
		return nil, err
	}

	err = DB.Select(&songs, "SELECT * FROM Song WHERE id IN VALUES( ? )", ids)

	return songs, err
}

func GetSongsPerPlaylist(playlistId string) ([]Song, error) {
	songs := []Song{}


	// SongInplaylist zooi
	err := DB.Select(&songs, "SELECT * From")
}