package config

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
)

var DbPool *sqlx.DB

func GetDbConnection() {
	var err error
	DbPool, err = sqlx.Connect("sqlite3", os.Getenv("DB_URL"))
	if err != nil {
		log.Println(err)
		panic("failed to connect to database")
	}
}

func MakeMigration() {
	DbPool.MustExec(`CREATE TABLE IF NOT EXISTS file (id varchar(255) primary key not null, source text not null);`)
}
