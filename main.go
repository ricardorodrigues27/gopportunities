package main

import (
	"github.com/ricardorodrigues27/gopportunities.git/config"
	"github.com/ricardorodrigues27/gopportunities.git/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")

	// Initialize Configs
	err := config.Init()

	if err != nil {
		logger.Error(err)
		return
	}

	// Initialize Router
	router.Initialize()
}
