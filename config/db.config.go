package config

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"os"
)

func GetDbConnection() (*sqlx.DB, error) {
	db, err := sqlx.Connect("sqlite3", os.Getenv("DB_URL"))

	return db, err
}

func MakeMigration(db *sqlx.DB) {
	db.MustExec(`CREATE TABLE IF NOT EXISTS file (id varchar(255) primary key not null, source text not null);`)
}
