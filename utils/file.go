package utils

import (
	"io"
	"os"
	"path/filepath"
)

type File interface {
	io.Reader
	WriteString(s string) (int, error)
	Close() error
	Name() string
	Sync() error
}

type OsFile struct {
	*os.File
	path  string
	close bool
}

type FileService interface {
	CreateTempFile() (File, error)
}

type FileOpenerService interface {
	FileService
	OpenFile(path string) (File, error)
}

type FileWriterService interface {
	FileService
}

type OsFileService struct {
}

func (o OsFileService) CreateTempFile() (File, error) {
	f, err := os.CreateTemp("", "")
	if err != nil {
		return nil, err
	}
	return File(OsFile{File: f}), nil
}

func (o OsFileService) OpenFile(path string) (File, error) {
	f, err := os.Open(filepath.Clean(path))
	if err != nil {
		return nil, err
	}
	return File(OsFile{File: f}), nil
}
