package storage

import "io"

type FileStorageClient interface {
	UploadFile(file io.Reader, filename string) (string, error)
	GetImage(filename string) ([]byte, error)
}
