package infrastructure

import (
	"github.com/labstack/echo"
	"interfaces/controllers"
)

func Init() {
	e := echo.New()
	sqlHandler := NewSqlHandler()
	albumController := controllers.NewAlbumController(sqlHandler)
	e.GET("/albums/:id", albumController.Show())
	e.Logger.Fatal(e.Start(":1323"))
}
