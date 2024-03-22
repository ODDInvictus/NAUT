package song_title

import (
	"log"
	"os"
	"strings"

	"github.com/nierot/invictus-radio/global_state"
	"github.com/spf13/viper"
	"github.com/zmb3/spotify/v2"
)



func Init() {
	state := global_state.GetGlobalState()

	state.AddCurrentSongListener(write)
}

func write(song *spotify.FullTrack) {
	// Write title - artists to a file on the disk
	// This is used to pass to icecast

	if song == nil || song.Name == "" {
		return
	}

	artists := []string{}
	for _, artist := range song.Artists {
		artists = append(artists, artist.Name)
	}

	str := song.Name + " - " + strings.Join(artists, ", ")

	// Write to file
	filename := viper.GetString("scheduler.song_title.filename")

	if filename == "" {
		log.Fatal("scheduler.song_title.filename is not set")
	}

	f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	// Truncate the file
	if err := os.Truncate(filename, 0); err != nil {
    log.Fatal("failed to truncate: %v", err)
}

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	// Write to file, overwriting the previous content
	_, err = f.WriteString(str + "\n")

	if err != nil {
		log.Fatal(err)
	}
}