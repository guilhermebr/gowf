package gowf

import (
	"fmt"
	"log"
	"net/http"
)

func (app *App) Run() {
	addr := HttpAddr

	if HttpPort != 0 {
		addr = fmt.Sprintf("%s:%d", HttpAddr, HttpPort)
	}

	s := &http.Server{
		Addr:    addr,
		Handler: app.Handlers,
	}

	err := s.ListenAndServe()

	if err != nil {
		log.Print(err)
	}
}
