package controllers

import (
	"github.com/labstack/echo"
	"net/http"
	"strconv"
	"interfaces/database"
	"usecase"
)

type AlbumController struct {
	Interactor usecase.AlbumInteractor
}

func NewAlbumController(sqlHandler database.SqlHandler) *AlbumController {
	return &AlbumController {
		Interactor: usecase.AlbumInteractor {
			AlbumRepository: &database.AlbumRepository {
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (controller *AlbumController) Show() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		album, err := controller.Interactor.AlbumById(id)
		if err != nil {

		}
		return c.JSON(http.StatusOK, album)
	}
}