package web

import "github.com/samverrall/hex-structure/internal/adapters/primary/web/users"

func (a *App) initAppRoutes() {
	a.fiber.Group("/users", users.CreateAccount(a.userAPI))
}
