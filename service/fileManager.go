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

func NewFileManager() *FileManager {
	return &FileManager{
		filepath.Base(".") + "/upload/",
		repository.NewFileRepository(),
	}
}

func (u FileManager) Save(file *multipart.FileHeader, authorizedMimeType []string, subDirectory string) (string, error) {
	filePath := u.basePath + subDirectory + u.formatFileName(file.Filename)
	fileOpen, err := file.Open()
	if err != nil {
		return "", errors.New("failed to open file")
	}
	defer fileOpen.Close()

	var buffer bytes.Buffer
	magicNumberBytes := make([]byte, 12)
	if _, err := fileOpen.Read(magicNumberBytes); err != nil {
		return "", err
	}

	magicNumber, err := u.GetMimeTypeFromMagicNumber(magicNumberBytes)
	isAuthorized := u.VerifyMimeType(magicNumber, authorizedMimeType)
	if !isAuthorized {
		return "", errors.New("file type unauthorized")
	}

	if _, err := io.Copy(&buffer, fileOpen); err != nil {
		return "", errors.New("file content is not readable")
	}

	err = os.WriteFile(filePath, buffer.Bytes(), 0644)
	if err != nil {
		return "", errors.New("failed to save file")
	}

	id, err := u.repository.SaveUploadFile(filePath)
	if err != nil {
		return "", err
	}
	return id, nil
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

func (u FileManager) GetMimeTypeFromMagicNumber(buffer []byte) (string, error) {
	if buffer[0] == 0xFF && buffer[1] == 0xD8 && buffer[2] == 0xFF {
		return "image/jpeg", nil
	} else if buffer[0] == 0x89 && buffer[1] == 0x50 && buffer[2] == 0x4E && buffer[3] == 0x47 {
		return "image/png", nil
	} else if string(buffer[0:4]) == "%PDF" {
		return "application/pdf", nil
	} else if string(buffer[0:4]) == "RIFF" && string(buffer[8:12]) == "WEBP" { // buffer[5:7] correspond to the file size
		return "image/webp", nil
	} else if string(buffer[0:3]) == "GIF" {
		return "image/gif", nil
	}
	return "", errors.New("unknown magic number type")
}

func (u FileManager) VerifyMimeType(mimeType string, authorizedMimeType []string) bool {
	for _, m := range authorizedMimeType {
		if m == mimeType {
			return true
		}
	}
	return false
}
