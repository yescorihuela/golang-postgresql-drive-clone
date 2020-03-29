package main

import (
	_ "github.com/yescorihuela/golang-postgresql-drive-clone/controllers"
	"github.com/yescorihuela/golang-postgresql-drive-clone/database"
	"github.com/yescorihuela/golang-postgresql-drive-clone/router"
	"github.com/yescorihuela/golang-postgresql-drive-clone/routes"
)

func main() {

	r := router.New()
	api := r.Group("/api")
	v1 := api.Group("/v1")
	d := database.New()
	database.AutoMigrate(d)
	routes.Register(v1)
	r.Logger.Fatal(r.Start(":8080"))
}
