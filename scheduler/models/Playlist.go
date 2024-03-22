package models

type Playlist struct {
	ID 									string `db:"id"`
	Name 								string `db:"name"`
	IsCurrentlyPlaying 	bool   `db:"isCurrentlyPlaying"`
	LastPlayed 					string `db:"lastPlayed"`
	CreatedAt 					string `db:"createdAt"`
	UpdatedAt 					string `db:"updatedAt"`
}

type PlaylistWithSongs struct {
	Playlist
	Songs []Song `json:"songs"`
}

func GetPlaylistByID(id string) (Playlist, error) {
	playlist := []Playlist{}
	err := DB.Select(&playlist, "SELECT * FROM Playlist WHERE id = ?", id)

	if len(playlist) == 0 {
		return Playlist{}, nil
	}

	return playlist[0], err
}

func GetPlaylists() ([]Playlist, error) {
	playlists := []Playlist{}
	// Select all playlists, sort by LastPlayed
	err := DB.Select(&playlists, "SELECT * FROM Playlist ORDER BY lastPlayed DESC")

	return playlists, err
}

// Update isCurrentlyPlaying and lastPlayed fields
func UpdatePlaylist(playlist Playlist) error {
	// Set all other playlists to not playing
	_, err := DB.Exec("UPDATE Playlist SET isCurrentlyPlaying = 0 WHERE id != ?", playlist.ID)

	if err != nil {
		return err
	}

	_, err = DB.Exec("UPDATE Playlist SET isCurrentlyPlaying = ?, lastPlayed = ? WHERE id = ?", playlist.IsCurrentlyPlaying, playlist.LastPlayed, playlist.ID)

	return err
}