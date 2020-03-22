package controllers

import (
	"net/http"
	"github.com/labstack/echo/v4"
)

func ListFiles(c echo.Context) error {
	e:= struct{
		Name string `json:"name"`
	}{
		"Yrvin Escorihuela",
	}

	return c.JSON(http.StatusOK, e)

}