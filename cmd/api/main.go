package main

import (
	"fmt"
	"github.com/eco-challenge/src/config"
	"github.com/eco-challenge/src/router"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	var err error
	if os.Getenv("DOCKER") != "true" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	config.GetDbConnection()
	if err = config.MakeMigration(); err != nil {
		fmt.Printf("Failed to execute migration : %s", err)
	}

	r := router.Provider()
	err = r.Run()
	if err != nil {
		panic("Router initialisation failed")
	}
}
