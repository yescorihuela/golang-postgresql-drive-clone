package repository

import (
	"github.com/yescorihuela/golang-postgresql-drive-clone/models"
)

func (fp *FileRepository) SaveNewFile(file *models.File) error{
	tx := fp.db.Begin()
	if err := fp.db.Create(file).Error; err != nil {
		return err
	}

	return tx.Commit().Error
}
