package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/yescorihuela/golang-postgresql-drive-clone/database"
)

type FileRepository struct{
	db *gorm.DB
}

func NewFileRepo() *FileRepository {
	return &FileRepository {
		db: database.New(),
	}
}