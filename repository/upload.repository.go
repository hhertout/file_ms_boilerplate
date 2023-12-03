package repository

import (
	"github.com/eco-challenge/config"
	uuid2 "github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type UploadRepository struct {
	db *sqlx.DB
}

func NewUploadRepository() UploadRepository {
	return UploadRepository{
		config.DbPool,
	}
}

func (u UploadRepository) SaveUploadFile(filePath string) error {
	uuid := uuid2.NewString()
	_, err := u.db.NamedExec("INSERT INTO file(id, source) VALUES (:uuid,:path)", map[string]interface{}{
		"uuid": uuid,
		"path": filePath,
	})
	if err != nil {
		return err
	}
	return nil
}

func (u UploadRepository) getUploadFile(id string) {

}
