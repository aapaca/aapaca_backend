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
	e.Logger.Fatal(e.Start(":1323"))
}
