package repository

import (
	"database/sql"
	"github.com/eco-challenge/src/config"
	uuid2 "github.com/google/uuid"
)

type FileRepository struct {
	db *sql.DB
}

func NewFileRepository() FileRepository {
	return FileRepository{
		config.DbPool,
	}
}

func (u FileRepository) SaveUploadFile(filePath string) (string, error) {
	uuid := uuid2.NewString()
	_, err := u.db.Exec("INSERT INTO file(id, source) VALUES ($1,$2)", uuid, filePath)
	if err != nil {
		return "", err
	}
	return uuid, nil
}

func (u FileRepository) GetUploadedFileById(id string) (string, error) {
	type Source struct {
		Source string `db:"source"`
	}

	source := Source{}
	rows, err := u.db.Query("SELECT source FROM file WHERE id=$1 LIMIT 1", id)

	if err != nil {
		return "", err
	}

	for rows.Next() {
		err = rows.Scan(&source.Source)
		if err != nil {
			return "", err
		}
	}

	return source.Source, err
}
