package repository

import (
	"github.com/eco-challenge/config"
	uuid2 "github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type FileRepository struct {
	db *sqlx.DB
}

func NewFileRepository() FileRepository {
	return FileRepository{
		config.DbPool,
	}
}

func (u FileRepository) SaveUploadFile(filePath string) error {
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

func (u FileRepository) GetUploadedFileById(id string) (string, error) {
	type Source struct {
		Source string `db:"source"`
	}

	source := Source{}
	rows, err := u.db.NamedQuery("SELECT source FROM file WHERE id=:id LIMIT 1", map[string]interface{}{
		"id": id,
	})

	if err != nil {
		return "", err
	}

	for rows.Next() {
		err = rows.StructScan(&source)
		if err != nil {
			return "", err
		}
	}

	return source.Source, err
}
