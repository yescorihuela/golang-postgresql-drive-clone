package responses

import (
	"time"
	"github.com/yescorihuela/golang-postgresql-drive-clone/models"
)

type fileResponse struct {
	File struct{
		ID uint `json:"id"`
		Name string `json:"name"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	} `json:"file"`
}

type filesResponse struct{
	Files []*fileResponse `json:"files"`
	FilesCount int `json:"files_count"`
}

func NewFileResponse(f *models.File) *fileResponse {
	r := new(fileResponse)
	r.File.ID = f.ID
	r.File.Name = f.Name
	r.File.CreatedAt = f.CreatedAt
	r.File.UpdatedAt = f.UpdatedAt
	return r
}

func NewFilesResponse(files []models.File) *filesResponse {
	r := new(filesResponse)
	r.Files = make([]*fileResponse, 0)
	r.FilesCount = len(files)
	for _, f := range files {
		fr := new(fileResponse)
		fr.File.ID = f.ID
		fr.File.Name = f.Name
		fr.File.CreatedAt = f.CreatedAt
		fr.File.UpdatedAt = f.UpdatedAt
		r.Files = append(r.Files, fr)
	}
	return r
}
