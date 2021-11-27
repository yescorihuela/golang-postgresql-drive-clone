package database

import (
	"fmt"
	"os"
	"strconv"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/yescorihuela/golang-postgresql-drive-clone/models"
)

func New() *gorm.DB {
	p := os.Getenv("DATABASE_PORT")
	port, err := strconv.Atoi(p)
	db, err := gorm.Open("postgres",
		fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
			os.Getenv("DATABASE_HOST"),
			port,
			os.Getenv("POSTGRES_USER"),
			os.Getenv("POSTGRES_PASSWORD"),
			os.Getenv("POSTGRES_DB"),
		))

	if err != nil {
		fmt.Println("storage err: ", err)
	}
	db.DB().SetMaxIdleConns(3)
	db.LogMode(true)
	return db
}

func Close(db *gorm.DB) {
	db.Close()
}

func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(
		&models.File{},
	)
}
