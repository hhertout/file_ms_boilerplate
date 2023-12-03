package service

import (
	"bytes"
	"errors"
	"github.com/eco-challenge/repository"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
)

type FileManager interface {
	Save(file *multipart.FileHeader) (string, error)
	Get(filename string) ([]byte, error)
}

type UploadManager struct {
	basePath   string
	repository repository.UploadRepository
}

func NewUploadManager() *UploadManager {
	return &UploadManager{
		filepath.Base(".") + "/upload/",
		repository.NewUploadRepository(),
	}
}

func (u UploadManager) Save(file *multipart.FileHeader) (string, error) {
	filePath := u.basePath + u.formatFileName(file.Filename)
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

func (u UploadManager) Get(filename string) ([]byte, error) {
	filePath := u.basePath + filename
	file, err := os.ReadFile(filePath)
	if err != nil {
		return nil, errors.New("file doesn't exist")
	}
	return file, nil
}

func (u UploadManager) GetBasePath() string {
	return u.basePath
}

func (u UploadManager) RegisterUpload() string {
	return ""
}

func (u UploadManager) formatFileName(filename string) string {
	return strings.Replace(filename, " ", "_", -1)
}
