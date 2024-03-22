package spotify

import (
	"context"
	"log"

	"github.com/nierot/invictus-radio/global_state"
	"github.com/zmb3/spotify/v2"
)

func GetCurrentlyPlaying() *spotify.FullTrack {
	state := global_state.GetGlobalState()

	if !state.SpotifyLoggedIn {
		log.Println("No spotify client found")
		return nil
	}

	cur, err := client.PlayerCurrentlyPlaying(context.Background())

	if err != nil {
		log.Println(err)
	}

	return cur.Item
}