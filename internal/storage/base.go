package storage

import "mime/multipart"

type FileStorageClient interface {
	UploadFile(file multipart.File, filename string) (string, error)
}
