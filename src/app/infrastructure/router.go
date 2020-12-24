package infrastructure

import (
	"interfaces/controllers"

	"github.com/labstack/echo"
)

func Init() {
	e := echo.New()
	sqlHandler := NewSqlHandler()

	albumController := controllers.NewAlbumController(sqlHandler)
	e.GET("/albums/:id", albumController.GetAlbum())
	e.GET("/artists/:id/albums", albumController.GetAlbumsByArtistId())

	songController := controllers.NewSongController(sqlHandler)
	e.GET("/songs/:id", songController.GetSong())
	e.GET("/artists/:id/songs", songController.GetAttendedSongs())
	e.GET("/albums/:id/songs", songController.GetSongsInAlbum())

	artistController := controllers.NewArtistController(sqlHandler)
	e.GET("/artists/:id", artistController.GetArtist())

	e.Logger.Fatal(e.Start(":1323"))
}
