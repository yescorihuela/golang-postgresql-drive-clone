package models

import (
	"github.com/jinzhu/gorm"
)

type File struct {
	gorm.Model
	Name         string `gorm:"unique_index;not null"`
	Md5Signature string `gorm:"unique_index;not null"`
	Location     string `gorm:"unique_index;not null"`
}

type User struct {
	gorm.Model
	Name            string `gorm:"not null"`
	Username        string `gorm:"unique_index;not null"`
	Email           string `gorm:"unique_index;not null"`
	Password        string `gorm:"not null"`
	ContractedQuota uint   `gorm:"not null;default:0"`
}

type UserFile struct {
	User   User
	UserID uint `gorm:"primary_key" sql:"type:int not null"`
	File   File
	FileID uint `gorm:"primary_key" sql:"type:int not null"`
}
