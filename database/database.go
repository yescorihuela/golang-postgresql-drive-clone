package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/yescorihuela/golang-postgresql-drive-clone/models"
)

func New() *gorm.DB {
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=drive-clone password=123456 sslmode=disable")
	if err != nil {
		fmt.Println("storage err: ", err)
	}
	db.DB().SetMaxIdleConns(3)
	db.LogMode(true)
	return db
}

func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(
		&models.File{},
	)
}