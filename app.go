package gowf

type App struct {
	Handlers *Router
}

func NewApp() *App {
	r := NewRouter()
	app := &App{Handlers: r}
	return app
}
