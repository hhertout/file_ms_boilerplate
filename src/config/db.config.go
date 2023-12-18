package config

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

var DbPool *sql.DB

func GetDbConnection() {
	var err error
	DbPool, err = sql.Open("sqlite3", os.Getenv("DB_URL"))
	if err != nil {
		log.Println(err)
		panic("failed to connect to database")
	}
}

func MakeMigration() error {
	_, err := DbPool.Exec(`CREATE TABLE IF NOT EXISTS file (
    id varchar(255) primary key not null, 
    source text not null);`,
	)
	if err != nil {
		return err
	}
	return nil
}
