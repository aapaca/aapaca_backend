package controller

import (
	"interfaces/repository"
	"interfaces/repository/rdb"
	"net/http"
	"strconv"
	"usecases/interactor"
	"usecases/usecase"

	"github.com/labstack/echo"
)

type AlbumController struct {
	Interactor usecase.AlbumUsecase
}

func NewAlbumController(sqlHandler rdb.SqlHandler) *AlbumController {
	return &AlbumController{
		Interactor: &interactor.AlbumInteractor{
			AlbumRepository: &repository.AlbumRepository{
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
			return c.JSON(http.StatusNotFound, APIError("Album Not Found"))
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
			return c.JSON(http.StatusNotFound, APIError("Albums Not Found"))
		}
		return c.JSON(http.StatusOK, albums)
	}
}
