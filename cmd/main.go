package main

import (
	"user-service/config"
	"user-service/internal/router"
)

func main() {
	config.LoadEnv()
	r := router.SetupRouter()
	r.Run(":8080")
}
