package infrastructure

import (
	"github.com/labstack/echo"
	"interfaces/controllers"
)

func Init() {
	e := echo.New()
	albumController := controllers.NewAlbumController(NewSqlHandler())
	e.GET("/album/:id", albumController.Show())
	e.Logger.Fatal(e.Start(":1323"))
}
