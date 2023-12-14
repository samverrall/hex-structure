package main

import (
	"log"

	"github.com/samverrall/hex-structure/internal/adapters/primary/web"
	"github.com/samverrall/hex-structure/internal/adapters/secondary/postgres"
	"github.com/samverrall/hex-structure/internal/core/services/users"
)

func main() {
	// Initialise secondary port implementations (Secondary adapters)
	userRepo, err := postgres.NewUserRepo() // <- this is swappable since its just a repo implementation
	if err != nil {
		log.Fatal("failed to init postgres user repo: %w", err)
	}

	// Initialise core service layer
	usersService := users.NewService(userRepo) // core business logic doesn't change.

	// Init primary (driving) adapter
	// this is swapple since we can spin up another primary adapter, and inject business logic
	srv := web.NewApp(usersService, web.WithPort(8000))

	srv.Run()
}
