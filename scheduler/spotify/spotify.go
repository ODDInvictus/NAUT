// This example demonstrates how to authenticate with Spotify using the authorization code flow.
// In order to run this example yourself, you'll need to:
//
//  1. Register an application at: https://developer.spotify.com/my-applications/
//     - Use "http://localhost:8080/callback" as the redirect URI
//  2. Set the SPOTIFY_ID environment variable to the client ID you got in step 1.
//  3. Set the SPOTIFY_SECRET environment variable to the client secret from step 1.
package spotify

// https://github.com/zmb3/spotify/blob/master/examples/authenticate/authcode/authenticate.go

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/nierot/invictus-radio/global_state"
	spotifyauth "github.com/zmb3/spotify/v2/auth"
	"golang.org/x/oauth2"
	"github.com/spf13/viper"

	"github.com/zmb3/spotify/v2"
)

// redirectURI is the OAuth redirect URI for the application.
// You must register an application at Spotify's developer portal
// and enter this value.
const redirectURI = "http://localhost:8080/callback"


var (
	auth  = spotifyauth.New(
		spotifyauth.WithRedirectURL(redirectURI), 
		spotifyauth.WithScopes(
			spotifyauth.ScopeUserReadPrivate,
			spotifyauth.ScopeUserReadCurrentlyPlaying,
			spotifyauth.ScopeUserReadPlaybackState,
			spotifyauth.ScopeUserModifyPlaybackState,
			))
	ch    = make(chan *spotify.Client)
	state = "abc123"
)

var token *oauth2.Token
var client *spotify.Client

func GetUrl() string {
	return auth.AuthURL(state)
}

func Init(e *echo.Echo) *spotify.Client {
	// first start an HTTP server
	e.GET("/callback", completeAuth)

	globalState := global_state.GetGlobalState()

	// TODO add client_id
	clientId := viper.GetString("spotify.client_id")
	clientSecret := viper.GetString("spotify.client_secret")

	if clientId == "" || clientSecret == "" {
		panic("spotify.client_id or spotify.client_secret not set in config")
	}

	log.Printf("spotify: %s %s\n", clientId, clientSecret)

	url := auth.AuthURL(state)
	log.Println("Please log in to Spotify by visiting the following page in your browser:", url)

	// wait for auth to complete
	client = <-ch
	
	// use the client to make calls that require authorization
	user, err := client.CurrentUser(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	globalState.AuthenticatedUser = user
	globalState.SetSpotifyLoggedIn(true)

	log.Println("You are logged in as:", user.ID)

	return client
}

func RefreshToken() {
	auth.RefreshToken(context.Background(), token)
}

func completeAuth(c echo.Context) error {
	r := c.Request()
	w := c.Response()

	tok, err := auth.Token(r.Context(), state, r)
	if err != nil {
		http.Error(w, "Couldn't get token", http.StatusForbidden)
		log.Fatal(err)
	}
	if st := r.FormValue("state"); st != state {
		http.NotFound(w, r)
		log.Fatalf("State mismatch: %s != %s\n", st, state)
	}

	token = tok

	// use the token to get an authenticated client
	client := spotify.New(auth.Client(r.Context(), tok))
	fmt.Fprintf(w, "Login Completed!")
	ch <- client

	return nil
}