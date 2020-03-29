package responses

import (
	"time"
	"github.com/yescorihuela/golang-postgresql-drive-clone/models"
)

type fileReponse struct {
	File struct {
		ID uint `json:"id"`
		Name string `json:"name"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	} `json:"file_record"`
}

func NewFileResponse(f *models.File) *fileReponse {

	r := new(fileReponse)
	r.File.ID = f.ID
	r.File.Name = f.Name
	r.File.CreatedAt = f.CreatedAt
	r.File.UpdatedAt = f.UpdatedAt
	return r
}