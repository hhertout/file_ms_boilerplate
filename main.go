package main

import (
	"github.com/eco-challenge/config"
	"github.com/eco-challenge/router"
)

func main() {
	err := config.LoadEnv()
	if err != nil {
		panic("Failed to load env")
	}

	r := router.Provider()
	err = r.Run()
	if err != nil {
		panic("Router initialisation failed")
	}
}
