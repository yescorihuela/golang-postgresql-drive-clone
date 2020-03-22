package main

import (
	"net/http"
	"github.com/labstack/echo/v4"
)

func main () {
	e := echo.New() // instance from echo framework
	s := struct{
		Status string `json:"status"`
	}{
		"API is fully functional...",
	}

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, s)
	})
	e.Logger.Fatal(e.Start(":8080"))
}
