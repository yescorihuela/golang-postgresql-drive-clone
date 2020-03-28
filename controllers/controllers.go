package controllers

import (
	"net/http"
	"github.com/labstack/echo/v4"
	"github.com/yescorihuela/golang-postgresql-drive-clone/repository"
	"github.com/yescorihuela/golang-postgresql-drive-clone/models"
)


func ListFiles(c echo.Context) error {
	e:= struct{
		Name string `json:"name"`
	}{
		"Yrvin Escorihuela",
	}
	return c.JSON(http.StatusOK, e)
}

func SaveNewFile(c echo.Context) error {
	file := new(models.File)
	c.Bind(&file)
	fileRepo := repository.NewFileRepo()
	if err := fileRepo.SaveNewFile(file); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, struct{
		Message string `json:"message"`
	}{
		file.Name,
	})
}