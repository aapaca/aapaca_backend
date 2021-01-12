package controllers

import (
	"interfaces/database"
	"interfaces/database/rdb"
	"net/http"
	"strconv"
	"usecase"

	"github.com/labstack/echo"
)

type AlbumController struct {
	Interactor usecase.AlbumInteractor
}

func NewAlbumController(sqlHandler rdb.SqlHandler) *AlbumController {
	return &AlbumController{
		Interactor: usecase.AlbumInteractor{
			AlbumRepository: &database.AlbumRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (controller *AlbumController) GetAlbum() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, APIError("Invalid Parameter"))
		}
		album, err := controller.Interactor.GetAlbum(id)
		if err != nil {
			return c.JSON(http.StatusBadRequest, APIError("Album Not Found"))
		}
		return c.JSON(http.StatusOK, album)
	}
}

func (controller *AlbumController) GetAlbumsByArtistId() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, APIError("Invalid Parameter"))
		}
		albums, err := controller.Interactor.GetAlbumsByArtistId(id)
		if err != nil {
			return c.JSON(http.StatusBadRequest, APIError("Albums not Found"))
		}
		return c.JSON(http.StatusOK, albums)
	}
}
