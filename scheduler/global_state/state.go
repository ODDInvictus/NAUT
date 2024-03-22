package global_state

import (
	"fyne.io/fyne/v2/data/binding"
	"github.com/zmb3/spotify/v2"
)

type GlobalState struct {
	SpotifyLoggedIn bool
	AuthenticatedUser *spotify.PrivateUser
	CurrentlyPlaying binding.Bool
	CurrentSong *spotify.FullTrack
	Tick int
	LastNewsBulletin *spotify.EpisodePage
}

var GlobalStateInstance = GlobalState{
	SpotifyLoggedIn: false,
	CurrentlyPlaying: binding.NewBool(),
	CurrentSong: nil,
	Tick: 0,
}

func GetGlobalState() *GlobalState {
	return &GlobalStateInstance
}

var loggedInListeners []func(loggedIn bool)

func (g *GlobalState) SetSpotifyLoggedIn(loggedIn bool) {
	g.SpotifyLoggedIn = loggedIn

	for _, listener := range loggedInListeners {
		listener(loggedIn)
	}
}

func (g *GlobalState) AddSpotifyLoggedInListener(listener func(loggedIn bool)) {
	loggedInListeners = append(loggedInListeners, listener)
}

var curSongListeners []func(song *spotify.FullTrack)

func (g *GlobalState) SetCurrentSong(song *spotify.FullTrack) {
	g.CurrentSong = song
	g.CurrentlyPlaying.Set(true)

	for _, listener := range curSongListeners {
		listener(song)
	}
}

func (g *GlobalState) AddCurrentSongListener(listener func(song *spotify.FullTrack)) {
	curSongListeners = append(curSongListeners, listener)
}

var tickListeners []func(tick int)

func (g *GlobalState) SetTick(tick int) {
	g.Tick = tick

	for _, listener := range tickListeners {
		listener(tick)
	}
}

func (g *GlobalState) AddTickListener(listener func(tick int)) {
	tickListeners = append(tickListeners, listener)
}