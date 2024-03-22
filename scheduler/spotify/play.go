package spotify

import (
	"context"
	"fmt"
	"log"

	"github.com/nierot/invictus-radio/global_state"
	"github.com/zmb3/spotify/v2"
)

func QueueByID(id spotify.ID) {
	state := global_state.GetGlobalState()

	fmt.Println("Queueing song with id", id)

	if !state.SpotifyLoggedIn {
		log.Println("No spotify client found")
		return
	}

	// Werkt niet met een show want hij doet spotify:track:id
	err := client.QueueSong(context.Background(), id)

	if err != nil {
		log.Println(err)
	}
}

func QueueShowByID(id spotify.ID) {
	state := global_state.GetGlobalState()

	fmt.Println("Queueing song with id", id)

	if !state.SpotifyLoggedIn {
		log.Println("No spotify client found")
		return
	}

	err := client.QueueEpisode(context.Background(), id)

	if err != nil {
		log.Println(err)
	}
}