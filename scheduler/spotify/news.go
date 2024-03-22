package spotify

import (
	"context"
	"log"
	"strconv"
	"time"

	"github.com/nierot/invictus-radio/global_state"
	"github.com/spf13/viper"
	"github.com/zmb3/spotify/v2"
)

func RunNews() {
	state := global_state.GetGlobalState()
	
	// First check the time
	// If it's time to play the news, play the news
	curTime := time.Now()

	// If it is before 7:00 and after 23:00, just skip
	if curTime.Hour() < 7 || curTime.Hour() > 23{
		return
	}

	// Check if the last news bulletin is still the same
	if state.LastNewsBulletin != nil {
		// check the title for the time
		name := state.LastNewsBulletin.Name
		// Name is of format Nieuwsbulletin 08:00
		// Get only the hours
		time := name[15:17]

		t, err := strconv.Atoi(time)
		if err != nil {
			log.Fatalln(err)
		}

		// If the current time is the same as the last news bulletin, return
		if curTime.Hour() == t {
			return
		}

		log.Println("last news message at: ", time)
	}

	// if it 5 minutes before a full hour, start looking for the next news bulletin
	if curTime.Minute() >= 55 {

	}

}

func GetNextNews() *spotify.EpisodePage {
	eps := queryNews()

	if eps == nil {
		return nil
	}

	// First eps is the latest
	ep := eps.Episodes[0]

	return &ep
}

func queryNews() *spotify.SimpleEpisodePage {
	state := global_state.GetGlobalState()

	if !state.SpotifyLoggedIn {
		return nil
	}

	newsId := viper.GetString("news.show_id")

	if newsId == "" {
		log.Fatal("news.show_id is not set")
	}

	eps, err := client.GetShowEpisodes(
		context.Background(), 
		newsId, 
		spotify.Market("NL"),
		spotify.Limit(1),
		spotify.Offset(0),
	)

	if err != nil {
		log.Println(err)
	}

	return eps
}