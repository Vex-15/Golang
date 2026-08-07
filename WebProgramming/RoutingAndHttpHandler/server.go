package main

import (
	"net/http"
	"time"
)

func (app *Application) serve() error {

	srv := http.Server{
		Addr:        ":8080",
		Handler:     app.routes(),
		ReadTimeout: 2 * time.Second,
	}
	return srv.ListenAndServe()
}
