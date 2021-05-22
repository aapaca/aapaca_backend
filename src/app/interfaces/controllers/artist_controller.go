package controllers

import (
	"interfaces/database"
	"interfaces/database/rdb"
	"net/http"
	"strconv"
	"usecase/interactor"

	"github.com/labstack/echo"
)

type ArtistController struct {
	Interactor interactor.ArtistInteractor
}

func NewArtistController(sqlHandler rdb.SqlHandler) *ArtistController {
	return &ArtistController{
		Interactor: interactor.ArtistInteractor{
			ArtistRepository: &database.ArtistRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (controller *ArtistController) GetArtist() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, APIError("Invalid Parameter"))
		}
		artist, err := controller.Interactor.GetArtist(id)
		if err != nil {
			return c.JSON(http.StatusBadRequest, APIError("Artist Not Found"))
		}
		return c.JSON(http.StatusOK, artist)
	}
}
