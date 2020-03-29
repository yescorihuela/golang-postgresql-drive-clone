package controllers

import (
	"net/http"
	"github.com/labstack/echo/v4"
	"github.com/yescorihuela/golang-postgresql-drive-clone/repository"
	"github.com/yescorihuela/golang-postgresql-drive-clone/models"
	"github.com/yescorihuela/golang-postgresql-drive-clone/responses"
)


func ListFiles(c echo.Context) error {

	fileRepo := repository.NewFileRepo()
	files, err := fileRepo.ListFiles()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	if files == nil {
		return c.JSON(http.StatusNotFound, struct{}{})
	}
	return c.JSON(http.StatusOK, responses.NewFilesResponse(files))
}

func SaveNewFile(c echo.Context) error {
	file := new(models.File)
	c.Bind(&file)

	fileRepo := repository.NewFileRepo()
	defer fileRepo.Close()

	if err := fileRepo.SaveNewFile(file); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, struct{
		Message string `json:"message"`
	}{
		file.Name,
	})
}