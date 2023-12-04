package main

import (
	"github.com/eco-challenge/config"
	"github.com/eco-challenge/router"
	"os"
)

func main() {
	var err error
	if os.Getenv("ENV") == "dev" {
		err = config.LoadEnv()
		if err != nil {
			panic(err)
		}
	}

	config.GetDbConnection()
	config.MakeMigration()

	r := router.Provider()
	err = r.Run()
	if err != nil {
		panic("Router initialisation failed")
	}
}
