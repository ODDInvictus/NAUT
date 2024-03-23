package main

import (
	"fmt"
	"log"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/nierot/invictus-radio/global_state"
	"github.com/nierot/invictus-radio/models"
	"github.com/nierot/invictus-radio/scheduler"
	"github.com/nierot/invictus-radio/song_title"
	"github.com/nierot/invictus-radio/spotify"
	"github.com/spf13/viper"
)

// Synchroniseer afspeellijsten
// verkrijg het nieuws
// soundboard
// Schrijf titel - artiest naar bestand
// Schrijf access token naar bestand
// Refresh token

var Queue *scheduler.Queue

func main() {
	log.Println("Starting Invictus Radio...")

	viper.SetConfigName("config.yaml")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("/etc/invictus-radio/")
	viper.AddConfigPath("$HOME/.invictus-radio")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("failed reading config")
		panic(err)
	}

	db := models.InitDB()
	defer db.Close()

	// Init webserver
	e := echo.New()
	e.Use(middleware.CORS())

	routes(e)

	// Init Queue
	Queue = scheduler.NewQueue()

	// Init spotify client
	go spotify.Init(e)

	// Init song title module
	go song_title.Init()

	// Init cron
	go cron()

	// Start the webserver
	e.Logger.Fatal(e.Start(":8000"))
}

func cron() {
	state := global_state.GetGlobalState()
	log.Println("Starting cron scheduler...")

	for range time.Tick(time.Second) {
		// Run every minute
		if state.Tick % 60 == 0 {
		}

		// Refresh token every 50 minutes
		if state.Tick % 3000 == 0 {
			go spotify.RefreshToken()
		}

		state.SetTick(state.Tick + 1)
	}
}