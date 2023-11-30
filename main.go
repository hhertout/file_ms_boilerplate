package main

import (
	"fmt"
	"github.com/eco-challenge/config"
	"github.com/eco-challenge/router"
)

func main() {
	err := config.LoadEnv()
	if err != nil {
		panic("Failed to load env")
	}

	dbPool, err := config.GetDbConnection()
	if err != nil {
		fmt.Println(err)
		panic("failed to connect to database")
	}
	config.MakeMigration(dbPool)

	r := router.Provider()
	err = r.Run()
	if err != nil {
		panic("Router initialisation failed")
	}
}
