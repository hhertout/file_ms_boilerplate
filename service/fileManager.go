package service

import (
	"errors"
	"github.com/eco-challenge/config"
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

	magicNumber, err := u.GetMimeTypeFromMagicNumber(file)
	isAuthorized := u.VerifyMimeType(magicNumber, authorizedMimeType)
	if !isAuthorized {
		return "", err
	}

	if err != nil {
		return "", errors.New("failed to open file")
	}
	defer fileOpen.Close()

	/*magicNumberBytes := make([]byte, 12)
	if _, err := fileOpen.Read(magicNumberBytes); err != nil {
		return "", err
	}

	magicNumber, err := u.GetMimeTypeFromMagicNumber(magicNumberBytes)
	isAuthorized := u.VerifyMimeType(magicNumber, authorizedMimeType)
	if !isAuthorized {
		return "", err
	}*/

	destinationFile, err := os.Create(filePath)
	if err != nil {
		return "", err
	}
	defer destinationFile.Close()

	if _, err := io.Copy(destinationFile, fileOpen); err != nil {
		return "", err
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

func (u FileManager) GetMimeTypeFromMagicNumber(file *multipart.FileHeader) (string, error) {
	fileOpen, err := file.Open()
	defer fileOpen.Close()

	if err != nil {
		return "", err
	}
	buffer := make([]byte, 12)
	if _, err := fileOpen.Read(buffer); err != nil {
		return "", err
	}

	// TODO build constants for mime type
	if buffer[0] == 0xFF && buffer[1] == 0xD8 && buffer[2] == 0xFF {
		return config.MIME_TYPE.Jpg, nil
	} else if buffer[0] == 0x89 && buffer[1] == 0x50 && buffer[2] == 0x4E && buffer[3] == 0x47 {
		return config.MIME_TYPE.Png, nil
	} else if string(buffer[0:4]) == "%PDF" {
		return config.MIME_TYPE.PDF, nil
	} else if string(buffer[0:4]) == "RIFF" && string(buffer[8:12]) == "WEBP" { // buffer[5:7] correspond to the file size
		return config.MIME_TYPE.Webp, nil
	} else if string(buffer[0:3]) == "GIF" {
		return config.MIME_TYPE.Gif, nil
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
