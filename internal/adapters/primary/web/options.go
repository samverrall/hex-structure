package web

type AppOption func(a *App)

// WithPort applies an optional port to the server.
func WithPort(port int) AppOption {
	return func(a *App) {
		a.port = port
	}
}
