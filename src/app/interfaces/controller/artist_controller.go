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

type ArtistController struct {
	Usecase usecase.ArtistUsecase
}

func NewArtistController(sqlHandler rdb.SqlHandler) *ArtistController {
	return &ArtistController{
		Usecase: &interactor.ArtistInteractor{
			ArtistRepository: &repository.ArtistRepository{
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

		artist, err := controller.Usecase.GetArtist(id)
		if err != nil {
			return c.JSON(http.StatusNotFound, APIError("Artist Not Found"))
		}
		return c.JSON(http.StatusOK, artist)
	}
}
