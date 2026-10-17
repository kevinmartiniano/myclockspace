package main

import (
	"myclockspace.kevinmartiniano.com/internal/app"
	"myclockspace.kevinmartiniano.com/internal/server"
)

func main() {
	router := app.NewRouter()

	server.NewServer(router).Up()
}
