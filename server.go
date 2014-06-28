package gowf

import (
	"fmt"
	"net/http"
)

func (app *App) Run() {
	addr := app.Config.HttpAddr

	if app.Config.HttpPort != 0 {
		addr = fmt.Sprintf("%s:%d", app.Config.HttpAddr, app.Config.HttpPort)
	}

	s := &http.Server{
		Addr:    addr,
		Handler: app.Handlers,
	}
	app.Logger.Printf("Running on %s", addr)

	err := s.ListenAndServe()

	if err != nil {
		app.Logger.Print(err)
	}

}
