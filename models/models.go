package models

import (
	"github.com/jinzhu/gorm"
)

type File struct {
	gorm.Model
	Name			string	`gorm:"unique_index:not null"`
}