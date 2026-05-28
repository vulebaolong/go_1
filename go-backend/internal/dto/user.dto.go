package dto

import "mime/multipart"

type UploadInput struct {
	FileHeader *multipart.FileHeader
	Folder     string
	Prefix     string
}

type UploadReturn struct {
	FileName string
	FullPath string
	Url      string
}
