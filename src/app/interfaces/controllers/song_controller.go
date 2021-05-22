package controllers

import (
	"interfaces/database"
	"interfaces/database/rdb"
	"net/http"
	"strconv"
	"usecases/interactor"
	"usecases/usecase"

	"github.com/labstack/echo"
)

type SongController struct {
	Interactor usecase.SongUsecase
}

func NewSongController(sqlHandler rdb.SqlHandler) *SongController {
	return &SongController{
		Interactor: &interactor.SongInteractor{
			SongRepository: &database.SongRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (controller *SongController) GetSong() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, APIError("Invalid Parameter"))
		}
		song, err := controller.Interactor.GetSong(id)
		if err != nil {
			return c.JSON(http.StatusBadRequest, APIError("Song Not Found"))
		}
		return c.JSON(http.StatusOK, song)
	}
}

func (controller *SongController) GetAttendedSongs() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, APIError("Invalid Parameter"))
		}
		songs, err := controller.Interactor.GetAttendedSongs(id)
		if err != nil {
			return c.JSON(http.StatusBadRequest, APIError("Songs Not Found"))
		}
		return c.JSON(http.StatusOK, songs)
	}
}

func (controller *SongController) GetSongsInAlbum() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, APIError("Invalid Parameter"))
		}
		songs, err := controller.Interactor.GetSongsInAlbum(id)
		if err != nil {
			return c.JSON(http.StatusBadRequest, APIError("Songs Not Found"))
		}
		return c.JSON(http.StatusOK, songs)
	}
}
