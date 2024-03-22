package main

import (
	"github.com/labstack/echo/v4"
	"github.com/nierot/invictus-radio/models"
)

type res struct {
	Message string 		
	Status int 				
	Data interface{} 	
}

func NotFound(c echo.Context, objectType string) error {
	return c.JSON(404, &res{
		Message: objectType + " niet gevonden",
		Status: 404,
		Data: nil,
	})
}

func InternalServerError(c echo.Context, err error) error {
	return c.JSON(500, &res{
		Message: err.Error(),
		Status: 500,
		Data: nil,
	})
}

func routes(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error {
		return c.JSON(200, &res{
			Message: "Invictus radio API!",
			Status: 200,
			Data: nil,
		})
	})

	e.GET("/queue", func(c echo.Context) error {
		if (Queue.Length() == 0) {
			return c.JSON(200, &res{
				Message: "queue is empty",
				Status: 200,
				Data: nil,
			})
		}

		return c.JSON(200, &res{
			Message: "Queue",
			Status: 200,
			Data: Queue.GetTracks(),
		})
	})

	e.GET("/playlist", func(c echo.Context) error {
		playlists, err := models.GetPlaylists()

		if err != nil {
			return InternalServerError(c, err)
		}

		return c.JSON(200, &res{
			Message: "Playlists",
			Status: 200,
			Data: playlists,
		})
	})

	e.GET("/playlist/:id", func(c echo.Context) error {
		id := c.Param("id")
		playlist, err := models.GetPlaylistByID(id)

		if err != nil {
			return InternalServerError(c, err)
		}

		if (playlist.ID == "") {
			return NotFound(c, "Afspeellijst")
		}

		return c.JSON(200, &res{
			Message: "Playlist",
			Status: 200,
			Data: playlist,
		})
	})

	e.PUT("/playlist/:id", func(c echo.Context) error {
		id := c.Param("id")
		playlist, err := models.GetPlaylistByID(id)

		if err != nil {
			return InternalServerError(c, err)
		}

		if (playlist.ID == "") {
			return NotFound(c, "Afspeellijst")
		}

		// Get request body to update playlist
		reqPlaylist := models.Playlist{}

		err = c.Bind(&reqPlaylist)

		if err != nil {
			return InternalServerError(c, err)
		}

		playlist.IsCurrentlyPlaying = reqPlaylist.IsCurrentlyPlaying
		playlist.LastPlayed = reqPlaylist.LastPlayed

		err = models.UpdatePlaylist(playlist)

		if err != nil {
			return InternalServerError(c, err)
		}

		return c.JSON(200, &res{
			Message: "Playlist updated",
			Status: 200,
			Data: playlist,
		})
	})
}
