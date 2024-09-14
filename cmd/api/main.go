package main

import (
	"github.com/StackItHQ/pes-devmegablaster/internal/config"
	"github.com/StackItHQ/pes-devmegablaster/internal/database"
)

func main() {
	config.Init()
	database.New()
}
