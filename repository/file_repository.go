package repository

import (
	"github.com/yescorihuela/golang-postgresql-drive-clone/models"
)

func (fp *FileRepository) SaveNewFile(file *models.File) error {
	tx := fp.db.Begin()
	if err := fp.db.Create(file).Error; err != nil {
		return err
	}
	return tx.Commit().Error
}

func (fp *FileRepository) ListFiles() ([]models.File, error) {
	var files []models.File
	if err := fp.db.Find(&files).Error; err != nil {
		return nil, err
	}
	return files, nil
}
