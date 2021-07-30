package test

import (
	"net/http"
	"net/http/httptest"

	"github.com/labstack/echo"
)

func CreateContextInstance(path string, paramName string, paramValue string) (*httptest.ResponseRecorder, echo.Context) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath(path)
	c.SetParamNames(paramName)
	c.SetParamValues(paramValue)
	return rec, c
}
