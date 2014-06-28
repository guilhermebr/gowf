package gowf

import (
	"log"
	"os"
)

type App struct {
	Config   *Config
	Handlers *Router
	Logger   *log.Logger
}

func NewApp() *App {
	return &App{
		Config:   InitConfig(),
		Logger:   log.New(os.Stdout, "", log.Ldate|log.Ltime),
		Handlers: NewRouter(),
	}
}
