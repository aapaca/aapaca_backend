package controllers

import (
	"interfaces/database"
	"net/http"
	"usecase"

	"github.com/labstack/echo"
)

type SongController struct {
	Interactor usecase.SongInteractor
}

func NewSongController(sqlHandler database.SqlHandler) *SongController {
	return &SongController{
		Interactor: usecase.SongInteractor{
			SongRepository: &database.SongRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (controller *SongController) GetSong() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strConv.Atoi(c.Param("id"))
		song, err := controller.Interactor.GetSong(id)
		if err != nil {
			return c.JSON(http.StatusBadRequest, APIError("Song Not Found"))
		}
		return c.JSON(http.StatusOK, song)
	}
}

func (controller *SongController) GetAttendedSongs() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strConv.Atoi(c.Param("id"))
		songs, err := controller.Interactor.GetAttendedSongs(id)
		if err != nil {
			return c.JSON(http.StatusBadRequest, APIError("Songs Not Found"))
		}
		return c.JSON(http.StatusOK, songs)
	}
}

func (controller *SongController) GetSongsInAlbum() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strConv.Atoi(c.Param("id"))
		songs, err := controller.Interactor.GetSongsInAlbum(id)
		if err != nil {
			return c.JSON(http.StatusBadRequest, APIError("Songs Not Found"))
		}
		return c.JSON(http.StatusOK, songs)
	}
}
