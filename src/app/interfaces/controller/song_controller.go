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

type SongController struct {
	Usecase usecase.SongUsecase
}

func NewSongController(sqlHandler rdb.SqlHandler) *SongController {
	return &SongController{
		Usecase: &interactor.SongInteractor{
			SongRepository: &repository.SongRepository{
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
		song, err := controller.Usecase.GetSong(id)
		if err != nil {
			return c.JSON(http.StatusNotFound, APIError("Song Not Found"))
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
		songs, err := controller.Usecase.GetAttendedSongs(id)
		if err != nil {
			return c.JSON(http.StatusNotFound, APIError("Songs Not Found"))
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
		songs, err := controller.Usecase.GetSongsInAlbum(id)
		if err != nil {
			return c.JSON(http.StatusNotFound, APIError("Songs Not Found"))
		}
		return c.JSON(http.StatusOK, songs)
	}
}
