package main

import (
	"github.com/Lucasvmarangoni/go-api/config"
	router "github.com/Lucasvmarangoni/go-api/routers"
)

var (
	logger *config.Logger
)

func main() {

	logger = config.GetLogger("main")

	err := config.Init()
	if err != nil {
		// panic(err)
		logger.Errorf("config initialization error: %v", err)
		return
	}
	router.Initialize()
}
