package web

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/samverrall/hex-structure/internal/core/services/users"
)

type App struct {
	fiber   *fiber.App
	userAPI users.API
	port    int
}

func NewApp(userAPI users.API, opts ...AppOption) *App {
	s := &App{
		fiber:   fiber.New(),
		userAPI: userAPI,
		port:    8000,
	}

	for _, applyOption := range opts {
		applyOption(s)
	}

	s.initAppRoutes()

	return s
}

func (a *App) Run() error {
	return a.fiber.Listen(fmt.Sprintf(":%d", a.port))
}
