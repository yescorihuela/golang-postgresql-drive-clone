package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/yescorihuela/golang-postgresql-drive-clone/controllers"
)

func Register(v1 *echo.Group) {
	files := v1.Group("/files")
	files.GET("/index", controllers.ListFiles)
	files.POST("/new", controllers.SaveNewFile)

}