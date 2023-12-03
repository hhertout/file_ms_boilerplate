package service

import (
	"bytes"
	"errors"
	"github.com/eco-challenge/repository"
	"github.com/google/uuid"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
)

type FileManager struct {
	basePath   string
	repository repository.FileRepository
}

func NewUploadManager() *FileManager {
	return &FileManager{
		filepath.Base(".") + "/upload/",
		repository.NewFileRepository(),
	}
}

func (u FileManager) Save(file *multipart.FileHeader, subDirectory string) (string, error) {
	filePath := u.basePath + subDirectory + u.formatFileName(file.Filename)
	fileOpen, err := file.Open()
	if err != nil {
		return "", errors.New("failed to open file")
	}
	defer fileOpen.Close()

	var buffer bytes.Buffer
	if _, err := io.Copy(&buffer, fileOpen); err != nil {
		return "", errors.New("file content is not readable")
	}

	err = os.WriteFile(filePath, buffer.Bytes(), 0644)
	if err != nil {
		return "", errors.New("failed to save file")
	}

	err = u.repository.SaveUploadFile(filePath)
	if err != nil {
		return "", err
	}
	return filePath, nil
}

func (u FileManager) Get(filename string) ([]byte, error) {
	filePath := u.basePath + filename
	file, err := os.ReadFile(filePath)
	if err != nil {
		return nil, errors.New("file doesn't exist")
	}
	return file, nil
}

func (u FileManager) GetBasePath(id string) (string, error) {
	path, err := u.repository.GetUploadedFileById(id)
	return path, err
}

func (u FileManager) formatFileName(filename string) string {
	id := uuid.New()
	e := filepath.Ext(filename)
	return id.String() + e
}
